# Practical Programming Analysis
##Assignment 2
####Timur, Andrew, Artem

####DataFlow Analysis for While programming language

####Language:
We took the language named While - simple imperative programming language designed to teach basic computer science and programming concepts. While is a procedural language that supports basic programming constructs such as loops, if-else statements, and variable assignment. It is considered the ancestor language of Pascal, and widely used to teach programming fundamentals in universities and colleges.

####DataFlow Analysis type:
We use dead code elimination. By performing a dead code elimination analysis, we can identify redundant code and eliminate it, freeing up resources and improving the program's overall performance. This analysis can also help reduce the risk of introducing bugs or vulnerabilities by removing any code that is no longer necessary. The elimination of dead code can also make a program easier to maintain and reduce the need for additional testing. As a result, the program can also become more scalable and easier to update as its feature set evolves.

#####Domain:
Variable names. 

#####Transfer functions:
print, :=, if, while

#####Confluence operator:
Meet, intersection
 
#####Initial values:
empty sets

#####Theoreticals properties:
Backward, must; our program tries to guarantee liveness of code: remove useless attributions, ifs and whiles
 

####Implementation details:
We follow the program from the bottom up, if we see a print of a variable, then we add it to the set of variables that need to be monitored.
If we see an expression in which one or more variables that we are monitoring changes, then we subtract from the set of variables that we are monitoring the set of variables changed in this expression, and add to the set of variables that we are monitoring the set of variables on which the result of the expression depends.
