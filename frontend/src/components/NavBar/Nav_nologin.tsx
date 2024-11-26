import { NavLink } from "react-router-dom";
import { Logout } from "../../services/https";
import Noti from "../../image/bell.svg"
import Premium from "../../image/diamond_636499.png"
import Logo from "../../image/washing-liquid (1).png"
import  "./Nav.css"
function Nav_nologin({page}: {  page:string}) {
    const isLogIn = localStorage.getItem("isLogin") === "true";
    return (
        <div >
             <>
               
               <div className="flex justify-between items-center  bg-white h-20 w-full rounded-b-3xl shadow-lg sticky px-10 top-0 ">
                <div className="flex justify-center items-center px-5 ">
                    <div>
                    <img src={Logo} alt="" className="w-10"/>    
                    </div>
                    <div className="text-2xl font-bold" onClick={Logout}>
                        verwash
                    </div>
        
                </div>
                 <div className="flex justify-between items-center border-none rounded-3xl bg-[#FBFBFB] py-2 px-5 gap-24">
                   <div >
                      <NavLink className={({ isActive }) => isActive || page == "Home" ? "Nav-font-active" : "Nav-font"} to="/Home" > Home</NavLink>

                   </div>
                   <div>
                      <NavLink className={({ isActive }) => isActive || page == "Order" ? "Nav-font-active" : "Nav-font"} to="/Order"> Order </NavLink>
                    </div>
                   <div>
                      <NavLink className={({ isActive }) => isActive || page == "Reservation" ? "Nav-font-active" : "Nav-font"} to="/Reserve"> Reserve </NavLink>
                    </div>
                    <div>
                      <NavLink className={({ isActive }) => isActive || page == "Report" ? "Nav-font-active" : "Nav-font"} to="/Report"> Report </NavLink>
                    </div>
                     
                 </div>
                 <div className="flex justify-start items-center">
                  <div className="flex justify-between items-center gap-5 border-none rounded-3xl bg-[#FBFBFB] py-2 px-5 mx-5">
                    <NavLink className="Nav-font" to="/Promotion" >
                    <div className="w-5 h-5">
                      <img src={Premium} alt="" className="w-5 h-5" />
                    </div>       
                    </NavLink>
                   <NavLink className="Nav-font m-0" to="/Notification">
                      <div id="noti">
                       <img src={Noti} alt=""  className="w-5 h-5 m-0"/>
                      </div>  
                   </NavLink>
                      
                 </div>
                    {isLogIn ? (
                        <NavLink className="Nav-font" to="/Profile">
                            Profile
                        </NavLink>
                    ) : (
                        <NavLink className="Nav-font" to="/login">
                            Login
                        </NavLink>
                    )}   
                 </div>
                 

               </div>
                   
                   
                </>
        </div>
    )
}
export default Nav_nologin;