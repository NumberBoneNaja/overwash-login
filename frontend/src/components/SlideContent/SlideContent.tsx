import { useEffect, useState } from "react"
import { PackageInterface } from "../../interface/IPackage"
import { GetAllPackage } from "../../services/https";
import "slick-carousel/slick/slick.css"; 
import "slick-carousel/slick/slick-theme.css";
import Slider from 'react-slick';
import "./Slide.css"
import Sneaker from "../../image/sneakers_2113286.png"
import { useNavigate } from "react-router-dom";

function SliceContent() {
    const [packageName, setPackageName] = useState<PackageInterface[]>([]);
    const navigate = useNavigate();

    async function getPackage() {
        const res = await GetAllPackage();
        setPackageName(res);
       
    }
    useEffect(() => {
        getPackage();
        
    }, []);

    useEffect(() => {
        console.log("p=",packageName); // จะได้ค่า packageName หลังจากที่มันถูกอัปเดต
    }, [packageName]); // จะทำงานเมื่อ packageName เปลี่ยนแปลง
    const settings = {
      dots: false,
      infinite: false,
      speed: 500,
      slidesToShow: 6,
      slidesToScroll: 5,
      arrow:true,
    };
    function GotoOrderDetail(id: number) {
      navigate(`/package-detail/${id}`);
    }
    return (
      <div className="content-service w-full h-auto px-10">
      <div className="content-serve  px-5">
        {packageName.length >= 6 ? ( // ตรวจสอบจำนวน packageName
          <Slider {...settings}>
            {packageName.map((item) => (
              <div
                className=" rounded-xl box-drop shadow-md h-[250px] w-[250px] px-5 py-5"
                key={item.ID}
              >
               
                <p>{item.ID}</p>
                <p>{item.PackageName}</p>
                <p>{item.Explain}</p>
              </div>
            ))}
          </Slider>
        ) : (
          // กรณีจำนวน packageName น้อยกว่า 6
          <div className="flex flex-wrap gap-4">
            {packageName.map((item) => (
              <div
                className="box1  bg-gradient-to-tl from-cyan-400 to-sky-400 rounded-xl box-drop shadow-md h-[250px] w-[250px] px-1 pt-5"
                key={item.ID}
                onClick={() => GotoOrderDetail(item.ID!)}
              > 
               <div className="detail">
                <div className="model w-full h-full flex justify-center">
                   <img src={Sneaker} alt="" className=" w-[100px] h-[100px]"/>  
                </div>
                <p className="font-bold text-center  text-white mt-5">{item.PackageName}</p>
                <p className=" text-white text-wrap">{item.Explain}</p>
                
               </div>
                
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
export default SliceContent