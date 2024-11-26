import Nav_bar from "../components/NavBar/Nav_bar";
import SliceContent from "../components/SlideContent/SlideContent";

function Order() {
    return (
        <div className="w-full h-full relative">
            <Nav_bar  page={"Order"}/>

            <div className="banner w-full h-80  flex justify-center my-5">
                <div className="content-in-banner w-[95%] h-full flex justify-center items-center 
                rounded-xl bg-[url('../../src/image/20994.jpg')]  bg-cover bg-center bg-no-repeat  ">
                    <div className="custombg h-full w-full bg-gray-600 bg-opacity-60  flex justify-center items-center  rounded-xl">
                         <h1 className="text-5xl font-medium font text-white">Laundry Order Package</h1>
                    </div>
                   
                </div>

            </div>
              <p className="text-2xl my-10 flex justify-center">PACKAGE การซัก</p>

              <SliceContent/>
            
        </div>
    );
}
export default Order;