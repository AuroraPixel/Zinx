package zface

type IRequest interface {
	//GetConnection 得到当前连接
	GetConnection() IConnection
	//GetData 得到请求请求数据
	GetData() []byte
}
