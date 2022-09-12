-- Given the CITY and COUNTRY tables, query the names of all cities where the CONTINENT is 'Africa'.


SELECT C.NAME FROM CITY C JOIN COUNTRY CT ON C.COUNTRYCODE=CT.CODE WHERE CT.CONTINENT = "Africa"