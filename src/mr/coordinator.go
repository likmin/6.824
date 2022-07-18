package mr

import "log"
import "net"
import "os"
import "net/rpc"
import "net/http"
import "strconv"
import "fmt"

type Coordinator struct {
	// Your definitions here.
	Sockname 	string
	workers   []int
	files     []string

}

// Your code here -- RPC handlers for the worker to call.
func (c *Coordinator) GetTask(args *MrWorker, task *MrTask) error {

	task.Filename = "likai" + strconv.Itoa(args.Sockname)

	fmt.Printf("listen from %v",args.Sockname)
	
	return nil
}
//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (c *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}


//
// start a thread that listens for RPCs from worker.go
//
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	c.Sockname = sockname
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
	ret := false

	// Your code here.


	return ret
}

//
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{}

	// Your code here.
	for i < len(files) {
		c.files = append(c.files, files[i])
	}
	
	for i < nReduce {
		c.workers = append(c.workers, i)
	}
	
	for i := range c.files {
		fmt.Printf("%v file need to be processed\n",i)
	}

	for i := range c.workers {
		fmt.Printf("%v worker\n",i)
	}
	c.server()
	return &c
}
