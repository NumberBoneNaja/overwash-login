import { SyntheticEvent, useState } from "react";
import Datepicker, { DateValueType } from "react-tailwindcss-datepicker";
function FormRegister() {
    const [regist, setRegist] = useState({
        username : '' , email : ''  , phone : '',gender : '',birthday:'', password : ''
     });
     function registchange(event: { target: { name: any; value: any; }; }){
        const {name, value} = event.target;
        setRegist(prevRegist => {
            return{
                ...prevRegist,
                [name] : value
            }
        })

     }
     function handlePhoneChange(event: { target: { value: string } }) {
        let value = event.target.value.replace(/\D/g, ""); // Remove non-numeric characters
        if (value.length > 3 && value.length <= 6) {
            value = `${value.slice(0, 3)}-${value.slice(3)}`;
        } else if (value.length > 6) {
            value = `${value.slice(0, 3)}-${value.slice(3, 6)}-${value.slice(6, 10)}`;
        }
        setRegist(prevRegist => ({
            ...prevRegist,
            phone: value,
        }));
    }
 
     function submit(e: SyntheticEvent){
        e.preventDefault();
        console.log(regist);
     }
    return (
        <div className="bg-white p-8 rounded-md shadow-md max-w-md mx-auto">
            
         
                <h1 className="text-xl font-bold mb-4">Register</h1>
            <div className="grid grid-cols-2 gap-4">
                <div className="mb-4">
                    <label htmlFor="username" className="block text-sm font-medium text-gray-700">Username</label>
                    <input
                        type="text"
                        id="username"
                        name="username"
                        className="border border-gray-300 p-2 w-full rounded"
                        onChange={registchange}
                        value={regist.username}
                        required
                    />
                </div>

                <div className="mb-4">
                    <label htmlFor="email" className="block text-sm font-medium text-gray-700">Email</label>
                    <input
                        type="email"
                        id="email"
                        name="email"
                        className="border border-gray-300 p-2 w-full rounded"
                        onChange={registchange}
                        value={regist.email}
                        required
                    />
                </div>
                <div>
                <label htmlFor="email" className="block text-sm font-medium text-gray-700">Email</label>
                    <input
                        type="date"
                        id="date"
                        name="date"
                        className="border border-gray-300 p-2 w-full rounded"
                        onChange={registchange}
                        value={regist.birthday}
                        required
                    />
                </div>
                <div className="mb-4">
                        <label htmlFor="phone" className="block text-sm font-medium text-gray-700">Phone number</label>
                        <input
                            type="tel"
                            id="phone"
                            name="phone"
                            placeholder="xxx-xxx-xxxx"
                            className="border border-gray-300 p-2 w-full rounded"
                            onChange={handlePhoneChange}
                            value={regist.phone}
                            maxLength={12} // Including dashes
                            required
                        />
                    </div>
                    <div>
                         <label htmlFor="gender">man</label>
                         <input type="radio" id="gender" name="gender" value="man" onChange={registchange} checked={regist.gender === "man"} ></input>
                         <label htmlFor="gender">woman</label>
                         <input type="radio" id="gender" name="gender" value="woman" onChange={registchange} checked={regist.gender === "woman"}></input>
                    </div>
                   

                <div className="mb-4">
                    <label htmlFor="password" className="block text-sm font-medium text-gray-700">Create Password</label>
                    <input
                        type="password"
                        id="password"
                        name="password"
                        className="border border-gray-300 p-2 w-full rounded"
                        onChange={registchange}
                        value={regist.password}
                        required
                    />
                </div>

                <div className="mb-4">
                    <label htmlFor="confirm-password" className="block text-sm font-medium text-gray-700">Confirm Password</label>
                    <input
                        type="password"
                        id="confirm-password"
                        name="confirmPassword"
                        className="border border-gray-300 p-2 w-full rounded"
                        required
                    />
                </div>
            </div>
                

                <button
                   onClick={submit}
                    type="submit"
                    className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition duration-200"
                >
                    Register
                </button>
         
        </div>
    );
}

export default FormRegister;
