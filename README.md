# Pick

Pick some data out of stdin and write it to stdout. 

### Example

to get the value dog from the following CSV line: 

```
echo "foo,bar,cat=dog" | pick -t csv -d 2 | pick -t qs -d cat
 > dog
```

or 

```
echo "foo=bar&bar=baz&cat=1" | pick -t qs -d bar 
> baz
```

if the CSV looks like this 

```
2016-01-02,cat,dog
2016-01-01,cat,dog
2016-01-02,dog,cat

```

you can count the unique values in the first column

```
cat data.csv | pick -d 1 | sort -n | uniq -c 
```