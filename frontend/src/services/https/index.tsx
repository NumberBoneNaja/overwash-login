

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

export async function GetAllPackage() {
  const token = localStorage.getItem("token");
  if (!token) {
    console.error("No token found. Please login.");
    return null; // ไม่มี token ส่งกลับ null
  }

  const requestOptions: RequestInit = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `${Bearer} ${token}`,
    },
    credentials: 'include' as RequestCredentials, // <--- Add this type annotation
  };

  try {
    const res = await fetch(`${ApiUrl}/GetAllPackage`, requestOptions);

    if (res.ok) {
      return await res.json();
    } if (res.status == 401) {
      console.log(token)
      window.location.href = '/login';
      return false
    }
    else {
      console.log(token)
      console.error("Failed to fetch data:", res.status, res.statusText);
      return null;
    }
  } catch (error) {
    console.error("Error fetching data:", error);
    return null;
  }
}

export async function GetClothByPackageID(id : any) {
  const token = localStorage.getItem("token");
  if (!token) {
    console.error("No token found. Please login.");
    return null; // ไม่มี token ส่งกลับ null
  }

  const requestOptions: RequestInit = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `${Bearer} ${token}`,
    },
    credentials: 'include' as RequestCredentials, // <--- Add this type annotation
  };

  try {
    const res = await fetch(`${ApiUrl}/GetClothByPackageID/${id}`, requestOptions);

    if (res.ok) {
      return await res.json();
    } if (res.status == 401) {
      console.log(token)
      window.location.href = '/login';
      return false
    }
    else {
      console.log(token)
      console.error("Failed to fetch data:", res.status, res.statusText);
      return null;
    }
  } catch (error) {
    console.error("Error fetching data:", error);
    return null;
  }
}

export async function GetAddOnByPackID(id : number) {
  const token = localStorage.getItem("token");
  if (!token) {
    console.error("No token found. Please login.");
    return null; // ไม่มี token ส่งกลับ null
  }

  const requestOptions: RequestInit = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `${Bearer} ${token}`,
    },
    credentials: 'include' as RequestCredentials, // <--- Add this type annotation
  };

  try {
    const res = await fetch(`${ApiUrl}/GetAddOnByPackageID/${id}`, requestOptions);

    if (res.ok) {
      return await res.json();
    } if (res.status == 401) {
      console.log(token)
      window.location.href = '/login';
      return false
    }
    else {
      console.log(token)
      console.error("Failed to fetch data:", res.status, res.statusText);
      return null;
    }
  } catch (error) {
    console.error("Error fetching data:", error);
    return null;
  }
}
export async function GetImageID(id : number) {
  const token = localStorage.getItem("token");
  if (!token) {
    console.error("No token found. Please login.");
    return null; // ไม่มี token ส่งกลับ null
  }

  const requestOptions: RequestInit = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `${Bearer} ${token}`,
    },
    credentials: 'include' as RequestCredentials, // <--- Add this type annotation
  };

  try {
    const res = await fetch(`${ApiUrl}/GetImagesByOrderID/${id}`, requestOptions);

    if (res.ok) {
      return await res.json();
    } if (res.status == 401) {
      console.log(token)
      window.location.href = '/login';
      return false
    }
    else {
      console.log(token)
      console.error("Failed to fetch data:", res.status, res.statusText);
      return null;
    }
  } catch (error) {
    console.error("Error fetching data:", error);
    return null;
  }
}
export async function UploadImageByOrderID(id: number, files: FileList) {
  const token = localStorage.getItem("token");
  if (!token) {
    console.error("No token found. Please login.");
    return null; // ไม่มี token ส่งกลับ null
  }

  const formData = new FormData();
  for (let i = 0; i < files.length; i++) {
    formData.append("images", files[i]);
  }

  const requestOptions: RequestInit = {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`, // ส่ง Authorization header
    },
    credentials: 'include' as RequestCredentials, // <--- Add this type annotation
    body: formData, // ใช้ FormData ตรงๆ
  };

  try {
    const res = await fetch(`${ApiUrl}/mutipleupload/${id}`, requestOptions);

    if (res.ok) {
      return await res.json(); // ถ้า upload สำเร็จ ให้ส่งผลลัพธ์จาก server
    } else if (res.status === 401) {
      console.log(token);
      window.location.href = '/login'; // ถ้า session หมดอายุหรือไม่มีสิทธิ์ให้ไปหน้า login
      return false;
    } else {
      console.log("Failed to fetch data:", res.status, res.statusText);
      return false;
    }
  } catch (error) {
    console.error("Error fetching data:", error);
    return null;
  }
}

 


