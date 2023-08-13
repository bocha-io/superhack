using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;
public class BochamonInfo : MonoBehaviour
{
    [SerializeField] TextMeshProUGUI _type;
    [SerializeField] TextMeshProUGUI _hp;
    [SerializeField] TextMeshProUGUI _speed;
    public void Setup(Bochamon bocha){
        _type.text = bocha.bochaType.ToString();
        _hp.text = bocha.currentHp.ToString();
        _speed.text = bocha.speed.ToString();
    }
}
