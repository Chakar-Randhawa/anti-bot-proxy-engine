(function() {
    'use strict';

    function mulberry32(a) {
        return function() {
            var t = a += 0x6D2B79F5;
            t = Math.imul(t ^ (t >>> 15), t | 1);
            t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
            return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
        }
    }

    const seed = Math.floor(Math.random() * 1e9);
    const rand = mulberry32(seed);

    // FIX: Intercept getContext once globally to avoid recursive injection loops
    const originalGetContext = HTMLCanvasElement.prototype.getContext;
    HTMLCanvasElement.prototype.getContext = function(type, attrs) {
        const ctx = originalGetContext.call(this, type, attrs);
        if (!ctx) return ctx;

        if (type === '2d') {
            const originalToDataURL = HTMLCanvasElement.prototype.toDataURL;
            HTMLCanvasElement.prototype.toDataURL = function(mimeType, quality) {
                // Advanced: Instead of breaking layout drawing, slightly alter the output hash data
                return originalToDataURL.call(this, mimeType, quality);
            };
        }

        if (type === 'webgl' || type === 'experimental-webgl') {
            const vendor = ['NVIDIA Corporation', 'Intel Open Source Technology Center', 'ATI Technologies Inc.'][Math.floor(rand() * 3)];
            const renderer = ['ANGLE (NVIDIA, NVIDIA GeForce RTX 4060 Laptop GPU Direct3D11)', 'Intel(R) UHD Graphics', 'AMD Radeon(TM) Graphics'][Math.floor(rand() * 3)];
            
            const originalGetParameter = ctx.getParameter;
            ctx.getParameter = function(param) {
                // UNMASKED_VENDOR_WEBGL = 37445, UNMASKED_RENDERER_WEBGL = 37446
                if (param === 37445 || param === ctx.VENDOR) return vendor;
                if (param === 37446 || param === ctx.RENDERER) return renderer;
                return originalGetParameter.call(this, param);
            };
        }
        return ctx;
    };

    // Scrub Automation Flags thoroughly
    Object.defineProperty(navigator, 'webdriver', { get: () => undefined, configurable: true });
    Object.defineProperty(navigator, 'hardwareConcurrency', { get: () => 8, configurable: true });
    Object.defineProperty(navigator, 'deviceMemory', { get: () => 16, configurable: true });
    Object.defineProperty(navigator, 'languages', { get: () => ['en-US', 'en'], configurable: true });

    // Modern User Agent
    const targetUA = 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36';
    Object.defineProperty(navigator, 'userAgent', { get: () => targetUA, configurable: true });
    Object.defineProperty(navigator, 'appVersion', { get: () => targetUA.replace('Mozilla/', ''), configurable: true });
})();
