grammar Grammar;

NUM : [0-9]+; 

t : b c
  ;

b : 'int'
  | 'float'
  ;

c : '[' n ']' c	
  | '[' n ']'
  ;

n : NUM
  ;
