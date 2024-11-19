import { Link, NavLink } from "react-router-dom";
import  "./Nav.css"
import { Logout } from "../../services/https";
import Nav_nologin from "./Nav_nologin";

function Nav_bar({ page}: {  page:string}) {
    const isLogIn = localStorage.getItem("isLogin") === "true";

    const userType:string = localStorage.getItem('user') ?? '';
    return (
        <div>
           {
             !isLogIn &&(
                <>
                   <Nav_nologin page={page}/>
                </>
             )
           }
           { isLogIn && userType == "admin" ?(
               <>
               <NavLink className={({ isActive }) => isActive ? "Nav-font-active" : "Nav-font"} to="/login"> Login</NavLink>
               <NavLink className={({ isActive }) => isActive ? "Nav-font-active" : "Nav-font"} to="/Home"> Home</NavLink>
               <NavLink className={({ isActive }) => isActive ? "Nav-font-active" : "Nav-font"} to="/Employee"> Employee</NavLink>
               <NavLink className={({ isActive }) => isActive ? "Nav-font-active" : "Nav-font"} to="#" onClick={Logout}> Logout</NavLink>
               </>
            ):(
                isLogIn && (
                    <>
                     <Nav_nologin page={page}/>
                    </>
                )
            )

           }

           
        </div>
    );
}
export default Nav_bar;