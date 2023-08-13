using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using TMPro;
using DG.Tweening;
public class MenuOption : MonoBehaviour
{
    [SerializeField] TextMeshProUGUI _name;
    public string menuName;
    public bool active;

    [SerializeField] protected Image _selection;
    [SerializeField] MenuOptionAction _action;

    public bool changeColor = true;


    public void Setup(string name){
        menuName = name;
        _name.text = name;
    }
    public void Execute(){
        _action.Execute();
    }

    public virtual void Select(bool s){
        _selection.gameObject.SetActive(s);
        if (s) {
            transform.localScale = new Vector3 (0.8f, 0.9f, 1);
            transform.DOScale(new Vector3(1,1,1), 0.1f);
            if (changeColor)
                _name.color = Color.white;
        }
        
        else _name.color = Color.black;
    }
}
