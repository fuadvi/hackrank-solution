-- Given the CITY and COUNTRY tables, query the names of all the continents (COUNTRY.Continent) and their respective average city populations (CITY.Population) rounded down to the nearest integer.

SELECT  CT.CONTINENT, FLOOR(AVG(C.POPULATION)) FROM CITY C JOIN COUNTRY CT ON C.COUNTRYCODE=CT.CODE GROUP BY 1 ORDER BY 2 ASC