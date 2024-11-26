import { useEffect, useState } from "react";
import { ClothType } from "../../interface/IClothType";
import { GetClothByPackageID } from "../../services/https";
import "./Pack.css"

interface CalculatorCloth {
    id :string,
    priceSelectedItemsChange: (selectedItems: number) => void;
    SelectedCloth:(selectedCloth:number[]) => void ;
}

function Package({id,priceSelectedItemsChange,SelectedCloth}: CalculatorCloth ){
    const [cloth, setCloth] = useState<ClothType[]>([]);
    const [counts, setCounts] = useState<{ [key: number]: number }>({});
    const [selectedItems, setSelectedItems] = useState<number[]>([]);
    const [totalPrice, setTotalPrice] = useState<number>(0);
    const [priceCloth, setPricecloth] = useState<{ [key: number]: number }>({});

   
    async function getCloth(id: number) {
         const res = await GetClothByPackageID(id);
         setCloth(res);

         const initialCounts = res.reduce((acc: { [key: number]: number }, item:ClothType) => {
            acc[item.ID!] = 0;
            return acc;
          }, {});
          setCounts(initialCounts);
         
    }
    useEffect(() => {
        getCloth(Number(id));
        console.log("cloth: ",cloth)
    },[]);
    useEffect(() => {
        const selectedClothDetails = cloth.filter((item) => selectedItems.includes(item.ID || 0));
        const calculatedTotalPrice = selectedClothDetails.reduce((total, item) => {
            const quantity = counts[item.ID || 0] || 0;
            return total + (item.Price || 0) * quantity;
        }, 0);
        
        setTotalPrice(calculatedTotalPrice);
        priceSelectedItemsChange(calculatedTotalPrice); // ส่งราคารวมกลับไปยังคอมโพเนนต์แม่
    }, [counts, selectedItems, cloth]);
    // useEffect(() => {
    //   const clothPrices = calculateClothPrices();
    //   setPricecloth(clothPrices); // อัปเดตราคาของผ้าแต่ละประเภท
    // }, [counts, cloth]);
   
    function plus(id: number) {
  
        setCounts((prevCounts) => ({
          ...prevCounts,
          [id]: (prevCounts[id] || 0) + 1,
        }));
      }
    function minus(id: number) {
    
        setCounts((prevCounts) => ({
          ...prevCounts,
          [id]: Math.max((prevCounts[id] || 0) - 1, 0), // ห้ามลดต่ำกว่า 0
        }));
       
      }
      const handleCheckboxChange = (id: number) => {
        setSelectedItems((prevSelectedItems) => {
          const updatedItems = prevSelectedItems.includes(id)
            ? prevSelectedItems.filter((itemId) => itemId !== id)
            : [...prevSelectedItems, id];//ใช้เก็บค่าที่ถูกเลือก

            SelectedCloth(updatedItems)
            setCounts((prevCounts) => {
              const updatedCounts = { ...prevCounts };
              if (updatedItems.includes(id)) {
                updatedCounts[id] = 1; // กำหนดค่าเริ่มต้นเป็น 1 เมื่อเลือก
              } else {
                delete updatedCounts[id]; // ลบค่าเมื่อยกเลิกการเลือก
              }
              return updatedCounts;
            });
          return updatedItems; // Add this return statement
        });
      };

      function calculateClothPrices(): { [key: number]: number } {
        const prices = cloth.reduce((acc: { [key: number]: number }, item) => {
          const quantity = counts[item.ID || 0] || 0; // จำนวนผ้าแต่ละประเภท
          acc[item.ID || 0] = (item.Price || 0) * quantity; // ราคาผ้าแต่ละประเภท = ราคา * จำนวน
          return acc;
        }, {});
        return prices;
      }
    
    return(
        <>
          <div className="detail-package w-full  ">
            <div className="w-full flex justify-center">
              <div className="table-cloth w-[95%] grid grid-cols-3 gap-7 px-5">
                <div>
                    <p className=" text-lg" >ชนิดผ้า</p>
                </div>
                <div>
                     <p  className="text-center text-lg">จำนวน</p>
                </div>
                <div>
                <p className="text-right text-lg">ราคา</p> 
                </div>
               
             </div>  
            </div>
            <div className="w-full">
                 {
                cloth.map((item,index)=>(
                    
                    <div key={index} className="w-full flex justify-center "  >
                        <div className={`table-cloth w-[95%] grid grid-cols-3 gap-7 px-5 my-2 h-[100px] rounded-xl shadow-md 
                        ${selectedItems.includes(item.ID || 0) ? "outline outline-1 outline-[#00BBF0] custom-shadow " : ""}`}>
                            <div className="flex items-center">
                                  
                                  <div>
                                  <div className="checkbox-wrapper-33">
                                    <label className="checkbox">
                                    <input className="checkbox__trigger visuallyhidden" type="checkbox" 
                                    checked={selectedItems.includes(item.ID || 0)} 
                                    onChange={() => handleCheckboxChange(item.ID || 0)}/>
                                     <span className="checkbox__symbol">
                                      <svg aria-hidden="true" className="icon-checkbox" width="28px" height="28px" viewBox="0 0 28 28" version="1" xmlns="http://www.w3.org/2000/svg">
                                        <path d="M4 14l8 7L24 7"></path>
                                         </svg>
                                            </span>
                                        
                                   </label>
                                    </div>
                                  </div>
                                 <p className="text-left text-xl">{item.TypeName}</p>
                                 
                            </div>
                            <div className="flex justify-center items-center border-solid ">
                              <div className="flex justify-center items-center bg-[#FBFBFB] rounded-l-lg  rounded-r-lg">
                                <button className="w-[25px]   rounded-l-lg text-white bg-[#24292E]"
                                onClick={() => minus(item.ID || 0)} disabled={counts[item.ID || 0] === 1 || !selectedItems.includes(item.ID || 0)}>-</button>
                                <p className="text-center mx-2">{counts[item.ID || 0] || 0}</p>  
                                <button className="w-[25px]  rounded-r-lg text-white bg-[#24292E]" disabled={ !selectedItems.includes(item.ID || 0)}
                                onClick={() => plus(item.ID || 0)}>+</button>
                              </div>
                                
                            </div>
                             <div className="flex justify-end items-center">
                              <div>
                                <p className="text-right">{item.Price} บาท</p>
                                {/* <p>{priceCloth[item.ID || 0]}</p>  */}
                              </div>
                               
                              
                             </div>
                           
                        
                        </div> 
                    </div>
                ))
               }   
                
               
          
           </div>  
          </div>
        </>
    )
}
export default Package 