

export const ApiUrl= "http://localhost:8000"
const Bearer = "Bearer"


export async function SignIn(data : any) {
    const requestOptions: RequestInit = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
        credentials: 'include', // ส่ง Cookie รวมกับการเรียก
    };
   
    
      let res = await fetch(`${ApiUrl}/login`, requestOptions)
        .then((res) => {
          if (res.status == 200) {
            console.log("congratulation")
            return res.json();
          }
          else if (res.status == 401) {
            console.log("fail")
      
            return false;
          }
           else ; {
            console.log("fail")
           
            return false
          }
        });
    
      return res;
}

export async function Logout() {
       
        const requestOptions: RequestInit = {
            method: "POST",
            credentials: 'include', // ส่ง cookie ที่เก็บ JWT ไปด้วย
        };
      
        const res = await fetch(`${ApiUrl}/logout`, requestOptions);
        if (res.ok) {
            console.log("Logged out successfully");
            localStorage.removeItem('token');
      
        // Remove the user from local storage
           localStorage.removeItem('user');
      
        // Remove the isLogin flag from local storage
           localStorage.removeItem('isLogin');
           localStorage.removeItem('id');
        // Redirect the user to the login page
          setTimeout(() => {
          window.location.href = '/login';
        }, 2000);
      
        }
   
        // Remove the token from local storage
        

}

//end Sign up