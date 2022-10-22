# Word Counter

A golang tool that counts words in a file.

## How to Install

Download the released binary

## How to Use

Assuming you have a text file called "myTextFile.txt", then:

<pre>
cat myTextFile.txt | ./WordCount
// *The number of words in the file will be displayed*
</pre>

If you want to count the number of lines instead, add the `l` flag:

<pre>
cat myTextFile.txt | ./WordCount -l
// *The number of lines in the file will be displayed*
</pre>
