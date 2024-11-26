interface CalculateProps {
    cloth: number;
    addon: number;
    selectedAddon:number[];
    selectedCloth:number[];
}

function Calculate({cloth, addon,selectedAddon,selectedCloth}: CalculateProps){
    function totalPrice(cloth: number, addon: number): number {
        
        return cloth + addon;
      }
      const Total = totalPrice(cloth,addon)
    return(
        <>
            <p>total is {Total}</p>
            <p>selec addon is  {selectedAddon}</p>
            <p>selec Cloth is  {selectedCloth}</p>
           
        </>
    )
}
export default Calculate