import { SyntheticEvent, useState } from "react";

import '../../src/index.css'; 
import LoginCom from "../components/Login/login";

function Login() {
     
    return (
        <div className="w-full h-screen bg-[url('../../src/image/20994.jpg')] bg-cover bg-center bg-no-repeat 
        bg-fixed bg-opacity-40 ">
            <div className="w-full h-full bg-black bg-opacity-50 fixed z-10 ">
                
            </div>
           <div className="flex justify-end w-full h-full fixed z-20">
             <LoginCom/>
           </div>
          

        
        </div>
    );
}
export default Login