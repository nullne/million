def infix2postfix(exp):
    postfix = ''
    stack = []
    def opOrder(op1,op2):
        order_dic = {'*':4,'$':5,'/':4,'+':3,'-':3}
        if op1 == '(' or op2 == '(':
            return False
        elif op2 == ')':
            return True
        else:
            if order_dic[op1] < order_dic[op2]:
                return False
            else:
                return True
    for s in exp:
        if s.isalpha():
            postfix += s
        else:
            while len(stack) != 0 and opOrder(stack[-1],s):
                op = stack.pop()
                postfix += op
            if len(stack) == 0 or s != ')':
                stack.append(s)
            else:
                top_op = stack.pop()
    if len(stack):
        postfix += ''.join(stack[::-1])
    return postfix
if __name__ == '__main__':
    for exp in ['A+B*C','(A+B)*C','((A-(B+C))*D)$(E+F)']:
        print infix2postfix(exp)
