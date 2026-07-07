**06/jul/2026**

a [friend](https://github.com/meet-dharmesh-gandhi) of mine asked me to implement a connection pool ( database connection ). Which i did. a followup that i got was to handle the situation in which if a connection given to a particular request, and the request goes into a inf loop how is the connection retrieved back ? 

We discussed the possibility of using a timer based approch in which i designed a min heap based approch where the top of the heap is the timer that will be going off the earliest. 


**07/jul/2026**

I decided to implement the timer as an project. but i wanted to make it a timer node. a node which handles the timer part and can be integrated into any other project or service. 

innitially i was going to implement this in node.js along with express.js but chatgpt convinced me to implement ths in go. 

 - ~ 11:00 : started the project ( no prior exp with go )
 - 18:18 : a working prototype with documentation ready for v1 release

