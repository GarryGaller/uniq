Uniq
======
~~~
~~~

Utility for searching for duplicate strings
---

```
>>uniq -help
uniq 1.0
Author: Garry G.

Usage of uniq:
uniq [-c|-d|-u] [-f num_fields] [-s skip_chars] [-w check_chars] [-range] [-color] [input] [output]
if input\output not specified, then stdin and stdout are used

  -c    Количество вхождений каждой строки
  -color
        Выделять использумый диапазон символов цветом
  -d    Вывести только повторяющиеся строки
  -f uint
        Игнорировать n полей разделенных пробелом с начала строки
  -i    Игнорировать регистр при сравнении строк
  -p string
        Количество строк в которых есть указанная подстрока
  -range
        Показать использумый диапазон символов как срез
  -s uint
        Игнорировать n символов с начала строки
  -u    Вывести только уникальные строки
  -w uint
        Проверять только n символов строки


```

Command line help
-----------------
***
**optional arguments:**


  * **-help**                  *Show this help message and exit.*
  * **-u**                     *Output only unique strings.*
  * **-d**                     *Output only lines that have repetitions.*
  * **-c**                     *Number of occurrences of each row*
  * **-p**                     *The number of rows in which there is a specified substring*  
  * **-f**                     *Skip N fields from the beginning of the string*
  * **-s**                     *Skip N characters from the beginning of the string.* 
  * **-w**                     *Check only n characters of the string.* 
  * **-color**                 *Highlight the used range of characters in color*  
  * **-range**                 *Show the used character range as a slice*

**unnamed arguments:**
input_file output_file
if not specified, then stdin and stdout are used 
~~~
~~~
EXAMPLES:  
=========

```
>>cat test.txt
AAA 0
aaa 0
ccc 1
ddd 7
eee 123
fff 246
ggg 249
hhh 369
iii 777
jjj 911
jjj 911
```


**by default print all lines excluding duplicates**
```
>>uniq test.txt
AAA 0
aaa 0
ccc 1
ddd 7
eee 123
fff 246
ggg 249
hhh 369
iii 777
jjj 911
```

**output only unique strings**
```
>>uniq -u test.txt
AAA 0
aaa 0
ccc 1
ddd 7
eee 123
fff 246
ggg 249
hhh 369
iii 777
```

**output only unique strings (ignoring case)**
```
>>>uniq -i -u test.txt
ccc 1
ddd 7
eee 123
fff 246
ggg 249
hhh 369
iii 777
```


**output only lines that have duplicates**
```
>>uniq -d test.txt
jjj 911
```

**output only lines that have duplicates (ignoring case)**
```
>>uniq -i -d test.txt
aaa 0
jjj 911
``` 


**number of occurrences of each row**
```
uniq -c test.txt
1 AAA 0
1 aaa 0
1 ccc 1
1 ddd 7
1 eee 123
1 fff 246
1 ggg 249
1 hhh 369
1 iii 777
2 jjj 911
```


**number of occurrences of each row (ignoring case)***
```
uniq -c test.txt
2 aaa 0
1 ccc 1
1 ddd 7
1 eee 123
1 fff 246
1 ggg 249
1 hhh 369
1 iii 777
2 jjj 911
```


**print the number of occurrences of the substring**
```
>>uniq -p=aaa test.txt
1 aaa
```

**print the number of occurrences of the substring (ignoring case)**
```
>>uniq -p=aaa -i test.txt
2 aaa
```


