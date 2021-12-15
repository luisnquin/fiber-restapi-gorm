# API RESTful with Fiber and GORM

### Main endpoints
#### Employees
<h5>GET, POST:</h5>

```http://127.0.0.1:8000/api/v1/employees```<br>

<h5>GET:</h5>

```http://127.0.0.1:8000/api/v1/employees/group/:id```<br>

<h5>GET, PUT, PATCH & DELETE:</h5>

```http://127.0.0.1:8000/api/v1/employees/:id```<br>

### Companies

<h5>GET, POST:</h5>

```http://127.0.0.1:8000/api/v1/companies```<br>

<h5>GET, PUT, PATCH & DELETE:</h5>

```http://127.0.0.1:8000/api/v1/companies/:id```<br>

<hr>
#### POST:
```http://127.0.0.1:8000/api/v1/companies```

```
{"name":"Mita","fundation_year":2010}
{"name":"Oozz","fundation_year":2008}
{"name":"Jatri","fundation_year":2006}
{"name":"Voomm","fundation_year":1999}
```

```http://127.0.0.1:8000/api/v1/employees```

```
{"fullname":"Chloette Robardet","age":42,"position":"VP Quality Control","company_code":1}
{"fullname":"Caprice Turbat","age":33,"position":"Teacher","company_code":1}
{"fullname":"Bobby Folland","age":42,"position":"Product Engineer","company_code":1}
{"fullname":"Aylmer Bentz","age":58,"position":"Help Desk Operator","company_code":2}
{"fullname":"Lindsay Rekes","age":52,"position":"VP Product Management","company_code":2}
{"fullname":"Tana Widocks","age":42,"position":"Research Associate","company_code":2}
{"fullname":"Darby Graver","age":45,"position":"Data Coordiator","company_code":3}
{"fullname":"Caitrin Spolton","age":5,"position":"Director of Sales","company_code":3}
{"fullname":"Kylen Critoph","age":56,"position":"Web Developer IV","company_code":3}
{"fullname":"Devin Cumes","age":28,"position":"Marketing Manager","company_code":4}
{"fullname":"Brigit McCray","age":57,"position":"General Manager","company_code":4}
{"fullname":"Shurlocke Pentony","age":19,"position":"Geologist I","company_code":4}
```

## Screenshots 

<img src="https://i.ibb.co/jJ3qqPS/carbon.png" alt="endpoints-routers">
