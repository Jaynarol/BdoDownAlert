## BdoDownAlert

###### Powered by [Jaynarol] in BdoTH family is [Noxia]

โปรแกรม BdoDownAlert ใช้สำหรับตรวจสอบการเชื่อมต่อของเกม เมื่อเกมหลุดการเชื่อมต่อ(Disconnect) หรือเกิดการเชื่อมต่อใหม่(Reconnect) โปรแกรมจะแจ้งเตือนให้ทราบทั้งการส่งเสียง เด้งป๊อปอัพ หรือแม้แต่ส่งข้อความเข้า Line
เหมาะสำหรับให้ทำงานช่วงที่ AFK ไม่อยู่หน้าจอ ดูหนัง นอน ไปทำงาน เพื่อให้โปรแกรมตรวจสอบการเชื่อมต่อเกมและแจ้งเตือนให้ทราบเมื่อมีปัญหา

โปรแกรม BdoDownAlert ไม่ใช่โปรแกรมโกง ไม่มีการ hack หรือไปยุ่งเกี่ยวกับการทำงานใดๆของเกม จึงไม่ส่งผลใดๆต่อเกมและไม่มีผลต่อการโดนแบน
หลักการทำงานของโปรแกรมคือตรวจสอบการเชื่อมต่อของอินเตอร์เน็ตด้วยคำสั่ง netstat ซึ่งเป็นคำสั่งพื้นฐานของคอมพิวเตอร์


#### Features
- ตรวจกับการ Reconnect (คือการหลุดแต่เกมยังเชื่อมต่ออัตโนมัติให้อยู่ เช่นเวลาตกปลาจะกลับมาตกต่อได้ แต่ถ้าแปรรูปหรือวิ่งม้าพอเข้ามาใหม่จะยืนเฉยๆ)
- ตรวจกับการ Disconnect (คือการหลุดออกจากเกมโดยสมบูรณ์เหมือนกดออกจากเกม)
- แจ้งเตือนด้วยการแสดง Popup อยู่บนสุดของทุกหน้าจอ ดูหนังเต็มจอ เล่นเกมอื่นอยู่ เห็นแน่นอน
- แจ้งเตือนด้วยการส่งเสียง สามารถปรับหน่วงเวลาการวนซ้ำเสียงแจ้งเตือนได้
- แจ้งเตือนด้วยการส่งข้อความเข้า Line (วิธีการเอา token อ่านได้ที่ https://bit.ly/2GQHJDB)
- สั่งให้โปรแกรม shutdown หรือ hibernate คอมได้หากเกมหลุดในช่วงเวลาที่ระบุ (เช่น ตี 1 ถึง 7 โมงเช้าหากเกมหลุดปิดคอมเลยไม่ต้องปลุก)
- มีการบันทึกการทำงานย้อนหลังเซฟเป็นไฟล์ไว้ให้ สามารถตรวจสอบได้ว่าหลุดตอนกี่โมง คอมปิดตอนไหน

![alt screenshot](https://i.imgur.com/QTf7pu1.png)
