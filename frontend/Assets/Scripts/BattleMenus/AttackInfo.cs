using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;
public class AttackInfo : MonoBehaviour
{
    [SerializeField] TextMeshProUGUI _type;
    [SerializeField] TextMeshProUGUI _power;
    public void Setup(Moves move){
        _type.text = move.moveType;
        _power.text = move.power.ToString();
    }
}
