using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using TMPro;
public class BochamonUI : MonoBehaviour
{
    [SerializeField] TextMeshProUGUI _name;
    [SerializeField] HPController _maxHp;
    public Bochamon bochamon;

    public void Setup(Bochamon bochi){
        bochamon = bochi;
        _name.text = bochi.bochaName;
        _maxHp.Setup(bochi.maxHp, bochi.currentHp);
    }

    public void ApplyDamage(int damage){
        _maxHp.ApplyDamage(damage);
    }

}
