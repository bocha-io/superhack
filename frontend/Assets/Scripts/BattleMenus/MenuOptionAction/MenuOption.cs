using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using TMPro;

public class MenuOption : MonoBehaviour
{
    [SerializeField] TextMeshProUGUI _name;
    public string menuName;
    public bool active;

    [SerializeField] Image _selection;
    [SerializeField] MenuOptionAction _action;


    public void Setup(string name){
        menuName = name;
        _name.text = name;
    }
    public void Execute(){
        _action.Execute();
    }

    public virtual void Select(bool s){
        _selection.gameObject.SetActive(s);
    }
}
