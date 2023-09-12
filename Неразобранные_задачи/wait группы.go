    go func(){
    
    var x1,x2 int
   
    //var resf1, resf2 int
        
        ch1:=make(chan int)
        ch2:=make(chan int)
        
        slice:=make([]int,n)
        
     wg:=new(sync.WaitGroup)
        
        for i:=0;i<=n;i++{ 
           
        wg.Add(1)
            
         x1=<-in1
         x2=<-in2   
       
    go func(i int,wg *sync.WaitGroup){
        defer wg.Done()
        defer close(ch1)
        defer close(ch2)
       
        go func(){
            ch1<-fn(x1)}()
        go func(){
            ch2<-fn(x2)}()
        
        slice[i]=<-ch1+<-ch2
        
    }(i,wg)
        
        }
        
        wg.Wait()
        
        for j:=0;j<n;j++{
        out<-slice[j]
    }
   }()
}