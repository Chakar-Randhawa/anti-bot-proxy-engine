package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	utls "github.com/refraction-networking/utls"
)

var (
	requestCount uint64
	successCount uint64
	blockCount   uint64
)

type MetricsMsg struct {
	Type       string `json:"type"` // "snapshot" or "log"
	Requests   uint64 `json:"requests"`
	Success    uint64 `json:"success"`
	Blocks     uint64 `json:"blocks"`
	Domain     string `json:"domain,omitempty"`
	Device     string `json:"device,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
	Timestamp  string `json:"timestamp,omitempty"`
}

// Thread-safe WebSocket Client Manager
type ClientManager struct {
	sync.RWMutex
	clients map[*websocket.Conn]bool
}

var cm = ClientManager{clients: make(map[*websocket.Conn]bool)}
var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Run periodic analytics stream
	go metricsBroadcaster(ctx)

	// Single Core Router on Port 8081 for UI and WebSockets
	http.HandleFunc("/", handleDashboard)
	http.HandleFunc("/ws", handleDashboardWS)

	// Start static UI and Monitoring server
	log.Println("🚀 Starting Web Dashboard Engine on http://localhost:8081")
	go func() {
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatalf("UI Engine Failure: %v", err)
		}
	}()

	// Reverse Proxy Middleware Engine on Port 8080
	proxyServer := &http.Server{
		Addr:         ":8080",
		Handler:      http.HandlerFunc(proxyHandler),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Println("🛡️ Advanced Anti-Bot Proxy Active and Listening on :8080")
	if err := proxyServer.ListenAndServe(); err != nil {
		log.Fatalf("Proxy Server Failure: %v", err)
	}
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dashboard.html")
}

func handleDashboardWS(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket Handshake Failure: %v", err)
		return
	}

	cm.Lock()
	cm.clients[conn] = true
	cm.Unlock()

	// Keep connection alive until client disconnects
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			cm.Lock()
			delete(cm.clients, conn)
			cm.Unlock()
			conn.Close()
			break
		}
	}
}

func metricsBroadcaster(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			snapshot := MetricsMsg{
				Type:     "snapshot",
				Requests: atomic.LoadUint64(&requestCount),
				Success:  atomic.LoadUint64(&successCount),
				Blocks:   atomic.LoadUint64(&blockCount),
			}
			broadcastMessage(snapshot)
		case <-ctx.Done():
			return
		}
	}
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// Handle raw CONNECT tunnels (HTTPS Passthrough optimization)
	if r.Method == http.MethodConnect {
		handleConnectTunnel(w, r)
		return
	}

	atomic.AddUint64(&requestCount, 1)
	req := r.Clone(r.Context())
	
	// FIX: Critical HTTP Core Error Resolution
	req.RequestURI = ""

	// Advanced: Modern JA4/TLS Fingerprint Engine Custom Dial Context
	dialTLSContext := func(ctx context.Context, network, addr string) (net.Conn, error) {
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}
		
		dialer := &net.Dialer{Timeout: 15 * time.Second}
		rawConn, err := dialer.DialContext(ctx, network, addr)
		if err != nil {
			return nil, err
		}

		// Inject Elite Browser Handshake profile (Chrome 120/JA4 Native Specs)
		config := &utls.Config{ServerName: host, InsecureSkipVerify: true}
		uconn := utls.UClient(rawConn, config, utls.HelloChrome_120)

		if err := uconn.Handshake(); err != nil {
			rawConn.Close()
			return nil, err
		}
		return uconn, nil
	}

	transport := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DialTLSContext:    dialTLSContext,
		ForceAttemptHTTP2: true,
		MaxIdleConns:      50,
		IdleConnTimeout:   30 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   20 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		atomic.AddUint64(&blockCount, 1)
		http.Error(w, "Proxy routing layer block or connection timeout.", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading secure upstream pipeline bytes.", http.StatusInternalServerError)
		return
	}

	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(strings.ToLower(contentType), "text/html") {
		strBody := string(bodyBytes)
		injectionCode := "<script>" + string(readLocalFile("injector.js")) + "</script>"
		
		// Streaming String Replacement injection strategy
		if idx := strings.Index(strings.ToLower(strBody), "<head>"); idx != -1 {
			idx += len("<head>")
			strBody = strBody[:idx] + injectionCode + strBody[idx:]
		} else {
			strBody = injectionCode + strBody
		}
		bodyBytes = []byte(strBody)
		resp.Header.Set("Content-Length", fmt.Sprintf("%d", len(bodyBytes)))
	}

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	w.Write(bodyBytes)

	// Evaluate Success metrics dynamically
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		atomic.AddUint64(&successCount, 1)
	} else {
		atomic.AddUint64(&blockCount, 1)
	}

	// Stream Dynamic Trace Logs to Web App Interface
	go broadcastMessage(MetricsMsg{
		Type:       "log",
		Domain:     r.URL.Host,
		Device:     "Chrome Emulated (JA4 Verified)",
		StatusCode: resp.StatusCode,
		Timestamp:  time.Now().Format("15:04:05"),
	})
}

func handleConnectTunnel(w http.ResponseWriter, r *http.Request) {
	destConn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		destConn.Close()
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		destConn.Close()
		return
	}

	// High-speed splice copy loops
	go func() {
		defer clientConn.Close()
		defer destConn.Close()
		io.Copy(destConn, clientConn)
	}()
	go func() {
		defer clientConn.Close()
		defer destConn.Close()
		io.Copy(clientConn, destConn)
	}()
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func readLocalFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		return []byte("console.log('Core payload injection error.');")
	}
	return data
}

func broadcastMessage(msg MetricsMsg) {
	cm.RLock()
	defer cm.RUnlock()
	for client := range cm.clients {
		_ = client.WriteJSON(msg)
	}
}
