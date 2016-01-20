import re 

class Parser:
	def Parse(self, expression):
		expr, num = [], []
		for number in re.findall(r'\d', expression):
			num.append(number)
		for exp in re.findall(r'\W', expression):
			expr.append(exp)

		return (("NUM", num), ("EXPR",expr))

def main():
	while True:
		expr = raw_input(">>> ")
		if expr:
			print Parser().Parse(expr)
		continue

if __name__ == "__main__":
	main()
	


