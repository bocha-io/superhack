using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;
public class AttackInfo : MonoBehaviour
{
    [SerializeField] TextMeshProUGUI _type;
    [SerializeField] TextMeshProUGUI _power;
    [SerializeField] TextMeshProUGUI _speed;
    public void Setup(Moves move){
        _type.text = move.moveType.ToString();
        _power.text = move.power.ToString();
        _speed.text = move.speed.ToString();
    }
}
