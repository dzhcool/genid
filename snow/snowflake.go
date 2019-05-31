package snow

// snowflake的结构如下(每部分用-分开):
// 0 - 0000000000 0000000000 0000000000 0000000000 0 - 00000 - 00000 - 000000000000
// 第一位为未使用，接下来的41位为毫秒级时间(41位的长度可以使用69年)，然后是5位datacenterId和5位workerId(10位的长度最多支持部署1024个节点） ，最后12位是毫秒内的计数（12位的计数顺序号支持每个节点每毫秒产生4096个ID序号）
// 一共加起来刚好64位，为一个Long型。(转换成字符串后长度最多19)

import (
	"log"
	"sync"
	"time"
)

type snowFlake struct {
	machineId     int64       // 机器码，占10位，0-1023
	incr          int64       // 自增编号，占12位，0-4095
	lastTimeStamp int64       // 上次生成时间，毫秒
	mux           *sync.Mutex // 互斥锁
}

var _snowFlake *snowFlake

func New() *snowFlake {
	if _snowFlake == nil {
		_snowFlake = new(snowFlake)
		_snowFlake.lastTimeStamp = _snowFlake.GetCurTimeStamp()
		_snowFlake.mux = new(sync.Mutex)
	}
	return _snowFlake
}

// 设置机器码
func (p *snowFlake) SetMachineId(mid int64) {
	p.mux.Lock()
	defer p.mux.Unlock()
	// 机器id左移12位，预留给自增序列号
	p.machineId = mid << 12
}

// 获取毫秒时间戳
func (p *snowFlake) GetCurTimeStamp() int64 {
	return time.Now().UnixNano() / 1000000
}

// 获取id
func (p *snowFlake) GenId() int64 {
	p.mux.Lock()
	defer p.mux.Unlock()

	curTimeStamp := p.GetCurTimeStamp()

	if curTimeStamp < p.lastTimeStamp {
		offset := p.lastTimeStamp - curTimeStamp

		if offset <= 10 {
			time.Sleep(time.Millisecond * time.Duration(offset*2))

			curTimeStamp := p.GetCurTimeStamp()
			if curTimeStamp < p.lastTimeStamp {
				log.Println("clock exception, retry faild last:", p.GetCurTimeStamp, " cur:", curTimeStamp)
				return 0
			}
		} else {
			log.Println("clock exception last:", p.GetCurTimeStamp, " cur:", curTimeStamp)
			return 0
		}
	}

	var id int64 = 0
	if curTimeStamp == p.lastTimeStamp {
		p.incr++
		if p.incr > 4095 {
			time.Sleep(time.Millisecond)
			curTimeStamp = p.GetCurTimeStamp()
			p.lastTimeStamp = curTimeStamp
			p.incr = 0
		}
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		rightBinValue <<= 22
		id = rightBinValue | p.machineId | p.incr
	}
	if curTimeStamp > p.lastTimeStamp {
		p.incr = 0
		p.lastTimeStamp = curTimeStamp
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		rightBinValue <<= 22
		id = rightBinValue | p.machineId | p.incr
	}
	return id
}
