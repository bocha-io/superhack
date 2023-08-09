using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class FightOption : MenuOption
{
    public Moves _move;
    [SerializeField] AttackInfo _attackInfo;
    [SerializeField] Image _moveType;
    public override void Select(bool s)
    {
        base.Select(s);
        if (s){
            _attackInfo.Setup(_move);
        }
    }

    public void Setup(Moves move){
        _move = move;
        _attackInfo.Setup(move);
        
        switch (move.moveType){
            case BochaType.Water:
                _moveType.color = new Color(0,0.63f, 0.86f);
                break;
            case BochaType.Fire:
                _moveType.color = new Color(0.86f, 0.17f, 0);
                break;
            case BochaType.Grass:
                _moveType.color = new Color(0.6f, 0.89f, 0.31f);
                break;
            default:
                _moveType.color = new Color(0.6f, 0.63f, 0.65f);
                break;
            
            
        }
        
        base.Setup(move.moveName);
    }
}
