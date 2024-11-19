import React from "react";

import { Navigate, Route, BrowserRouter as Router, Routes } from "react-router-dom";

import Login from "./pages/login";
import Home from "./pages/Home";
import Employee from "./pages/Employee";
import Member from "./pages/Member";
import ProtectedRoute from "./components/ProtectRoute/ProtectRoute";
import ErrorPage from "./pages/Error";
import './index.css'; // Import Tailwind CSS
import Rigister from "./pages/Register";
import Order from "./pages/Order";
import Reserve from "./pages/Reserve";
import Reports from "./pages/Report";
import Profile from "./pages/Profile";
import Notifications from "./pages/Notification";
import Promotion from "./pages/Promotiom";

const App: React.FC = () => {
  const isLogIn = localStorage.getItem("isLogin") === "true";

    const userType:string = localStorage.getItem('user') ?? '';
  return (
  <>
     
     
      <Routes>
       {/* unauthorizes */}
       {!isLogIn && (
            <>
              <Route path="/login" element={<Login />} />
              <Route path="/" element={<Home />} />
              <Route path="/Employee" element={<Login />} />
              <Route path="/Home" element={<Home />} />
              <Route path="/Rigister" element={<Rigister />} />
              <Route path="/Order" element={<Login />} />
              {/* <Route path="/Member" element={<Login />} /> */}
            </>
          )}
        {/* หลังเข้าระบบ */}
        <Route element={<ProtectedRoute />}>
            <Route path="/" element={<Home />} />
            <Route path="/Order" element={<Order />} />
            <Route path="/Employee" element={<Employee />} />
            <Route path="/Home" element={<Home />} />
            <Route path="/Member" element={<Member />} />
            <Route path="/Reserve" element={<Reserve />} /> 
            <Route path="/Report" element={<Reports />} /> 
            <Route path="/Profile" element={<Profile />} />
            <Route path="/Notification" element={<Notifications />} />
            <Route path="/Promotion" element={<Promotion />} />
        </Route>
        
        <Route path="*" element={<ErrorPage />} />
        

     </Routes>
  </>
    

  );

};


export default App;