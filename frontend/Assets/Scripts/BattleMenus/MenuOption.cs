using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class MenuOption : MonoBehaviour
{
    public string menuName;
    public bool active;
    [SerializeField] Image _selection;
    [SerializeField] MenuOptionAction _action;
    public void Execute(){
        _action.Execute();
    }

    public void Select(bool s){
        _selection.gameObject.SetActive(s);
    }
}
