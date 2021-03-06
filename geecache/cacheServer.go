package geecache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type CacheServer struct {
	addr     string
	basePath string
	tm       *TicketManager
}

func NewCacheServer(addr, basePath string, tm *TicketManager) *CacheServer {
	return &CacheServer{
		addr:     addr,
		basePath: basePath,
		tm:       tm,
	}
}

// ServeHTTP handle all http requests
func (p *CacheServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Log("%s %s", r.Method, r.URL.Path)
	key := r.URL.Query().Get("key")
	count := r.URL.Query().Get("count")
	startNo := r.URL.Query().Get("startNo")
	endNo := r.URL.Query().Get("endNo")
	userID := r.URL.Query().Get("userID")
	passangerIDs := strings.Split(r.URL.Query().Get("passangerIDs"), "&")
	fmt.Println(userID)
	fmt.Println(passangerIDs)

	res, err := p.tm.BuyTickets(key, count, startNo, endNo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	//消息队列生产者,消息包括同步库存和生成订单
	//1 同步库存，把修改的库存同步到数据库

	//2 生成订单，对每个passangerID和res中的其中一个座位号，生成订单
}

// Log info with server name
func (p *CacheServer) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.addr, fmt.Sprintf(format, v...))
}
