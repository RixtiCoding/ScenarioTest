
# Data Processing Pipeline Scenario
## Building and running the project
### REQUIRED:
#### GO 1.18+
#### Docker 


#### STEP 1 
**Download the project zip file from github**

#### STEP 2
- **cd into the project from your preffered terminal**
- cd into /api/cmd and run ```go build``` -> This will build the api exec
- cd back into ScenarioTest-Master(the main project file) and now run ```go build``` -> this will build the main exec

#### STEP 3
run ```docker run --name=rediboard -p 6379:6379 redis``` into your terminal to instantiate a redis database connection
OPTIONAL: if you wish to have a gui for the redis database interaction, use [Redis GUI](https://redis.com/redis-enterprise/redis-insight/#insight-form)

# VERY IMPORTANT !
as the project uses hard-coded messages and sample data, in order for the program to see the paths of the .json files you will have to run the execs from the same directory as they are located in.
EXAMPLE -> after you built the api exec, while you're still in ScenarioTest-master/api/cmd
run ```./cmd```
EXAMPLE -> after you built the main exec, while you're still in ScenarioTest-master, run ```./MessageFilter```



# - Design Decisions
 **CLI** - why have i used a CLI for this scenario?
Using a cli made the project much more focused around the core filter logic. This comes with downsides also, such as no continuous message filtering. As this is a test scenario, with an emphasis on the logic part, i found this approach to be perfect, minimizing the code that is written and making it a little bit easier to understand, but a little bit more messy.
**Testin approach** 
- i've written tests for most of the important functions
- the api does not have testing as it is too simple to be worth doing it.
- I did not wanted to not use any popular testing tehniques or third party libraries for simplicity sake.
**The approval service**
- I've tried to do an extremely short representation on how it would look like on the cli. it required no human interaction as it approves/rejects on a random criteria.
 
