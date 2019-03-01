package mystack

type Mystack struct {
	array []interface{}
	size int
	count int
}

func NewStack(n int) (*Mystack,bool) {
	if n <= 0 {
		return nil,false
	}
	return &Mystack{
		array:make([]interface{},n),
		size:n,
		count:0,
	},true
}

func (this *Mystack)Push(v interface{}) bool {
	if this.size <= 0 {
		return false
	}
	if this.count >= this.size {
		//fmt.Printf("stack is full! can not push!")
		//return false
		//支持动态扩容
		this.array = append(this.array,v)
		this.size = cap(this.array)
		this.count = len(this.array)
		return true
	}
	this.count++
	this.array[this.count-1] = v
	return true
}

func (this *Mystack)Pop() (v interface{}) {
	if this.count <= 0 {
		return nil
	}
	v = this.array[this.count-1]
	this.array[this.count-1] = nil
	this.count--
	return v
}

func (this *Mystack)Size() int {
	return this.size
}

func (this *Mystack)Count() int {
	return this.count
}