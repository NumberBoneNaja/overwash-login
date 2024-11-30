import Nav_bar from "../components/NavBar/Nav_bar";
import { useParams } from "react-router-dom";
import Banner from "../components/Banner/Banner";
import Package from "../components/PackageContent/Package";

import AddOn from "../components/AddOn/AddOn";
import { useEffect, useState } from "react";
import Calculate from "../components/Calculate/calculate";
import { ImageInterface } from "../interface/IImage";
import {  GetImageID, UploadImageByOrderID } from "../services/https";
import { ClothSelectedDetail } from "../components/PackageContent/Package";
import { AddOnInterface } from "../interface/IAddOn";



function PackageDetail() {
    const { id } = useParams<{ id: string }>();
    const [totalcloth, setTotalcloth] = useState<number>(0);
    const [totalAddOn, setTotalAddOn] = useState<number>(0);
    const [selectedAddOn,setSelectedAddOn] = useState<number[]>([]);
    const [selectedCloth,setSelectedCloth] = useState<number[]>([]);
    const [selectedClothDetail, setSelectedClothDetail] = useState<ClothSelectedDetail[]>([]);
    const [addOn,setAddon] = useState<AddOnInterface[]>([]);
    
    
    
    // async function GetImage() {
    //     const res = await GetImageID(1);
    //     setImage(res);
    // }
    
    // const handleUpload = async (id: number, files: FileList | null) => {
    //     if (!files || files.length === 0) {
    //         console.error("No files selected.");
    //         return;
    //     }

    //     const result = await UploadImageByOrderID(id, files);

    //     // ตรวจสอบผลลัพธ์จากการอัปโหลด
    //     if (!result) {
    //         console.error("Upload failed or no result returned.");
    //         return;
    //     }

    //     if (result.error) {
    //         console.error("Error:", result.error); // หากมี error ให้แสดงผล
    //     } else {
    //         console.log("Upload successful:", result);
    //         setUploadStatus("Upload successful");
    //     }
    // };
    // const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    //     // Check if files were selected
    //     const files = event.target.files;
    //     if (files) {
    //       // Handle the selected files
    //       setSelectedFiles(files);
    //     }
    //   };
    useEffect(()=>{
        // console.log("detail cloth: ",selectedClothDetail)
        // console.log("total price: ",totalcloth)
        // console.log("total price addon: ",totalAddOn)
        // console.log("selected cloth: ",selectedCloth)
        // console.log("selected addon: ",selectedAddOn)
        // console.log("image4: ",image)
    },[totalcloth,totalAddOn,selectedCloth,selectedAddOn]
    )
    
    


    return (
        <div className="bg-[#FBFBFB]">
            <Nav_bar  page={"Order"}/>
            <Banner/>
            <p className="text-2xl my-10 text-center">ประเภทเสื้อ</p>
            <Package id= {id as string} priceSelectedItemsChange={setTotalcloth} SelectedCloth={setSelectedCloth} detailClothselected={setSelectedClothDetail} />
            <p className="text-2xl my-10 text-center">บริการเสริม</p>
            <AddOn id={id as string} priceSelectedAddOnChange ={setTotalAddOn} SelectedAddOn={setSelectedAddOn} AddOnItem={setAddon} />
            <Calculate cloth ={totalcloth} addon = {totalAddOn} selectedAddon={selectedAddOn}  selectedCloth={selectedCloth}  detailsCloth={selectedClothDetail} 
            id={id || ''} addonItem = {addOn}/>
            {/* <div>
            <input type="file" multiple onChange={handleFileChange} />
            <button onClick={() => handleUpload(1, selectedFiles)}>Upload</button>
            </div> */}
            {/* <div>
                {
                    image.map((item)=>(
                        <div key={item.ID}>
                            <img src={item.ImagePath} alt="" />
                        </div>
                    ))
                }
            </div> */}
        </div>
    );
}

export default PackageDetail;