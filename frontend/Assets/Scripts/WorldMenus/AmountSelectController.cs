using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;

public class AmountSelectController : MonoBehaviour
{
    public int maxCoins;
    public int amount;
    [SerializeField] TextMeshProUGUI _amountText;
    
    public void Update(){
        if(Input.GetKeyDown(KeyCode.D)){
            PressRight();
        }
        if(Input.GetKeyDown(KeyCode.A)){
            PressLeft();
        }

        if(Input.GetKeyDown(KeyCode.F)){
            SendCoins();
        }
    }

    public void PressRight(){
        amount +=1;
        if (amount > maxCoins){
            amount=maxCoins;
        }
        _amountText.text = amount.ToString();
    }

    public void PressLeft()
    {
        amount -=1;
        if (amount < 0){
            amount=0;
        }
        _amountText.text = amount.ToString();
    }

    public void SendCoins()
    {

    }

}
