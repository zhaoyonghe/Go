import sys
import os

result = []
with open('question.txt','r') as f:
	for line in f:
		result.append(line.strip('\n'))

for line in result:
	if not os.path.exists(line):
		os.mkdir(line)
	with open(line + "\\solution1.go",'w') as f:
		pass