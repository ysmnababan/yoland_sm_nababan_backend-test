SELECT m.id, m.name, p.status as 'pendidikan terakhir', m.time_create , p.time_create as 'time_update'
FROM `murid` as m  
join `pendidikan` as p
on p.id_murid = m.id
WHERE p.id in (
	select pId from ( 
		select p.id_murid, MAX(p.id) as pId 
		from `pendidikan` as p
		GROUP by p.id_murid
	) as pTemp
);