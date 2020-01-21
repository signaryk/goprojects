# goprojects
Projects I'm working on while learning the Go programming language

# Projects
- Synchronous Multithreading

## Synchronous Multithreading
My first project with Go was to implement a solution the classic reader/writer category of problems. Coming from a C++ background, 
I started with using mutexes to protect the shared data and control the activity of reader and writer threads. Once with this solution completed,
I then took the previous implementation and modified it to instead use Go's native channels to pass the data between readers and writers. 
For this type of problem, using channels proved to be a much easier implementation. 
