import { SyntheticEvent, useState } from "react";
import { Form, Link, useNavigate } from "react-router-dom";
import { SignIn } from "../../services/https";
import { jwtDecode } from "jwt-decode";
import { message } from "antd";

function LoginCom() {
    const [messageApi, contextHolder] = message.useMessage();
    const success = () => {
        messageApi.open({
          type: 'success',
          content: 'login success',
        });
      };
      const error = () => {
        messageApi.open({
          type: 'error',
          content: 'Email or Password is incorrect',
        });
      };
     const [sign, setSign] = useState({
        email : '' , password : ''
     });
     

     function signchange(event: { target: { name: any; value: any; }; }){
        const {name, value} = event.target;
        setSign(prevSign => {
            return{
                ...prevSign,
                [name] : value
            }
        })

     }
     async function submit(e: SyntheticEvent){
        e.preventDefault();
        
        const check = await SignIn(sign);
        console.log(check);
        if (check === false) {
            error();
        }
        else{
        success();
        localStorage.setItem('token',check.token);
        localStorage.setItem('user',check.position);
        localStorage.setItem('isLogin', `${true}`);
        const user_type = localStorage.getItem('user');
        if (user_type === 'admin') {
            setTimeout(() => {
                window.location.href = '/Employee'; 
            }, 2000);
             
        }else if(user_type === 'user') {
            setTimeout(() => {
                window.location.href = '/Member'; 
            }, 2000);
        }
        try {
            const token = localStorage.getItem('token');
            const infotoken:any = jwtDecode(token||'');
            localStorage.setItem('id', infotoken.id);
        } catch (error) {
            console.log(error);
        }
        }
       
     }

    return (
        <div >
            {contextHolder}
        
        <div className="w-80 bg-white h-full px-2 rounded-l-3xl drop-shadow-xl z-10" style={{width:"500px" }}>
            
        <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
           
        <div className="sm:mx-auto sm:w-full sm:max-w-sm">
        <h2 className=" text-center text-4xl font-bold tracking-tight text-pink-500  ">
            Welcome!!
          </h2> 
          <h2 className="mt-10 text-center text-2xl/9 font-bold tracking-tight text-gray-900">
            Sign in
          </h2>
        </div>

        <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
          <form action="#" method="POST" className="space-y-6">
            <div>
              <label htmlFor="email" className="block text-sm/6 font-medium text-gray-900">
                Email address
              </label>
              <div className="mt-2">
                <input
                  id="email"
                  name="email"
                  type="email"
                  value={sign.email}
                  onChange={signchange}
                  required
                  placeholder="Enter your email"
                  autoComplete="email"
                  className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset sm:text-sm/6
                 px-2"
                />
              </div>
            </div>

            <div>
              <div className="flex items-center justify-between">
                <label htmlFor="password" className="block text-sm/6 font-medium text-gray-900">
                  Password
                </label>
              
              </div>
              <div >
                <input
                  type="password"
                  value={sign.password}
                  name = "password"
                  onChange={signchange}
                  autoComplete="current-password"
                  placeholder="Password"
                  className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset  sm:text-sm/6
                
                  px-2 "
                />
             
              
              </div>
            </div>

            <div>
              <button
                type="submit"
                className="flex w-full justify-center rounded-md bg-black px-3 py-1.5 text-sm/6 font-semibold text-white shadow-sm hover:bg-gray-800 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
              onClick={submit}>
                Sign in
              </button>
            </div>
          </form>

          <p className="mt-10 text-center text-sm/6 text-gray-500">
            Not a member?{' '}
            <Link to ="/Register" className="font-semibold  hover:text-pink-300 text-pink-500">
              Rigister
            </Link>
          </p>
        </div>
      </div>
            

        
        </div>
        </div>
            
    );
}
export default LoginCom