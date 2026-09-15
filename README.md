# 🛡️ Advanced Anti-Bot Mitigation Bypass Engine (A.B.M.E.)
### *Next-Generation Low-Latency Network Proxy & Stealth Browser Telemetry Instrumentation Suite*

---

## 💡 Project Overview & Reality Context
In modern software engineering, web scraping, data aggregation, and automated testing are heavily restricted by automated anti-bot security networks (such as **Cloudflare Turnstile, DataDome, and Akamai**). Traditional headless browser frameworks (Selenium, Puppeteer, Playwright) or basic HTTP headers-only spoofing are instantly flagged, causing connection bans and data extraction failures.

This project is a high-performance **Localized Forward Proxy Engine and Injection Framework** written from scratch in Go (Golang). It intercepts outbound automated browser traffic and dynamically mutates low-level browser fingerprinters, structural JA4 TLS connection profiles, and runtime JavaScript runtime variables to make automated scripts indistinguishable from genuine premium consumer devices—running completely on your local machine with **Zero Cloud Infrastructure Fees**.

---

## 🎯 Strategic Real-World Impact & Business Value

- **Massive Enterprise Cost Reduction ($0 Running Costs):** Eliminates the need for expensive premium commercial residential proxy rotation services or un-capped API bypass networks that cost companies thousands of dollars monthly.
- **Ultra-Lightweight & Low Spec Optimized:** Built using explicit memory pooling and continuous binary streaming logic, ensuring the framework runs seamlessly with a minimal footprint on standard consumer laptops (even under heavy concurrent traffic).
- **Asynchronous Live Telemetry Dashboard:** Includes a clean, responsive single-file dark management interface streaming live metrics (Throughput/s, Success vs Block ratios, Dynamic Trace logs) over continuous WebSockets.

---

## 🏗️ System Architecture & Data Flow

[ Headless / Automated Browser Client ]│▼ (Forwards Outbound HTTP/HTTPS Traffic)[ Go Proxy Middleware Engine (:8080) ]│├─► [ uTLS Transport Layer ] ───► Emulates Chrome 120 / JA4 Handshaking├─► [ Script Injection Node ] ──► Injecting injector.js into DOM on-the-fly│▼ (Streams Event Analytics Concurrently)[ WebSocket Channel Core (:8081/ws) ] ───► [ Live Control Panel UI ] ───► Interactive Charts & Logs│▼[ Upstream Target Server (Cloudflare Protected) ] ───► [ Bypass Successful - HTTP 200 OK ]

---

## 🛠️ Deep Technical & Architectural Implementations

### 1. Dual-Port Multithreaded Routing Kernel (`main.go`)
- **Memory Optimization:** Engineered using sequential processing memory buffers to handle high network volumes without memory leaks or file descriptor exhaustions.
- **Precision DOM Injection:** Utilizes a single-pass streaming logic string matcher that intercepts remote server HTML responses and instantly forces our stealth payload script right below the opening `<head>` block before compilation occurs.
- **Thread-Safe Sockets Synchronization:** Protects active client pointer maps using a strict `sync.RWMutex` read-write lock pattern to safely handle overlapping data transmissions without triggering runtime race conditions.

### 2. Low-Level Browser Telemetry Overriding (`injector.js`)
- **Automation Flag Deletion:** Scrubs structural browser tracking attributes (such as `navigator.webdriver`) and securely freezes runtime modifications by locking property configurations using `Object.defineProperty`.
- **Hardware Profile Spoofing:** Overwrites local processing parameters (`hardwareConcurrency`) and physical ram indicators (`deviceMemory`) to match authentic premium workstation models.
- **Advanced WebGL Overriding:** Intercepts unmasked hardware graphics configurations (`VENDOR`, `RENDERER`) to safely supply fake, deterministic consumer GPU values.

### 3. Asynchronous Visualization Panel (`dashboard.html`)
- **Reactive WebSocket Core:** Processes distinct structural data snapshots and text-based raw trace logs independently over highly persistent full-duplex pipelines.
- **Low Overhead Rendering:** Implements non-blocking `Chart.js` algorithms designed to smoothly plot moving line charts and doughnut ratios at 1-second ticks with virtual zero rendering overhead.

---

## 🚀 Step-by-Step Local Deployment Manual (A to Z)

Setting up and executing the engine on a local architecture is simple. Follow the sequence below:

### Step 1: Verification of Prerequisites
Ensure the **Go (Golang)** compiler workspace environment is running on your system. 
```bash
# Verify Go is successfully mapped to your system path
go version
```

### Step 2: Clone the Application Workspace
Open your system terminal/command prompt and clone this public repository:
```bash
git clone https://github.com
cd anti-bot-proxy
```

### Step 3: Synchronize System Dependencies
Pull the external optimized networking modules and Gorilla socket structures:
```bash
go mod tidy
```

### Step 4: Run the Core Architecture Suite
Execute the multi-threaded compilation script to boot both microservices:
```bash
go run main.go
```
Upon successful initialization, your terminal will output the following operational status logs:
- `🚀 Starting Web Dashboard Engine on http://localhost:8081`
- `🛡️ Advanced Anti-Bot Proxy Active and Listening on :8080`

### Step 5: Verify Live Traffic Manipulation
1. Launch your browser and navigate to the monitoring control link: **`http://localhost:8081`**.
2. Configure your automated script, HTTP client, or manual browser to route external internet traffic through your active local proxy engine address: **`127.0.0.1:8080`**.
3. Access any secure web endpoint. The proxy core will instantly intercept the session, spoof underlying fingerprints, and stream real-time throughput spikes and dynamic status logs straight to your visible dashboard charts.
