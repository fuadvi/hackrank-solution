-- Given the CITY and COUNTRY tables, query the sum of the populations of all cities where the CONTINENT is 'Asia'.

SELECT SUM(C.POPULATION) FROM CITY C JOIN COUNTRY CT ON C.COUNTRYCODE=CT.CODE WHERE CT.CONTINENT = "Asia"