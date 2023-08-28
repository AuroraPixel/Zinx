package znet

import "Zinx/zface"

type Request struct {
	conn zface.IConnection
	data []byte
}

// GetConnection 得到当前连接
func (r *Request) GetConnection() zface.IConnection {
	return r.conn
}

// GetData 得到请求请求数据
func (r *Request) GetData() []byte {
	return r.data
}
