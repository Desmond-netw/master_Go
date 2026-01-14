func main

// Building lookup maps

func manufacturerMap (manufacturer []Manufacturer) map[int]string {
	manuMap := make(map[int]string)
for , manuf range manufacturer {
	manuMap[manuf.ID] = manuf.Name
}
return manuMap
}