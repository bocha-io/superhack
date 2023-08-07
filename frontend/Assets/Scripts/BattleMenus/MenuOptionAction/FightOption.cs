using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class FightOption : MenuOption
{
    public Moves _move;
    [SerializeField] AttackInfo _attackInfo;
    public override void Select(bool s)
    {
        base.Select(s);
        if (s){
            _attackInfo.Setup(_move);
            Debug.Log("Setup Move");
        }
    }

    public void Setup(Moves move){
        _move = move;
        _attackInfo.Setup(move);
        base.Setup(move.moveName);
    }
}
