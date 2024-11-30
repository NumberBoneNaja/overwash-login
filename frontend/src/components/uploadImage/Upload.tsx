import { useEffect, useState } from "react";
import { ImageInterface } from "../../interface/IImage";
import { GetImageID, UploadImageByOrderID } from "../../services/https";

function UploadImage({OrderID} : {OrderID:string}){
    const [image,setImage] = useState<ImageInterface[]>([]);
    const [selectedFiles, setSelectedFiles] = useState<FileList | null>(null);
    const [uploadStatus, setUploadStatus] = useState<string>("");
    async function GetImage() {
        const res = await GetImageID(1);
        setImage(res);
    }
    const handleUpload = async (id: number, files: FileList | null) => {
        if (!files || files.length === 0) {
            console.error("No files selected.");
            return;
        }

        const result = await UploadImageByOrderID(id, files);

        // ตรวจสอบผลลัพธ์จากการอัปโหลด
        if (!result) {
            console.error("Upload failed or no result returned.");
            return;
        }

        if (result.error) {
            console.error("Error:", result.error); // หากมี error ให้แสดงผล
        } else {
            console.log("Upload successful:", result);
            setUploadStatus("Upload successful");
        }
    };
    const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        // Check if files were selected
        const files = event.target.files;
        if (files) {
          // Handle the selected files
          setSelectedFiles(files);
        }
      };
    useEffect(()=>{
         console.log(image)
       },[image]
    )
    useEffect(()=>{
        GetImage()
    
    },[]
    )

    return(
        <>
        <div>
            <input type="file" multiple onChange={handleFileChange} />
            <button onClick={() => handleUpload(Number(OrderID), selectedFiles)}>Upload</button>
            </div>
        </>
    )
}
export default UploadImage