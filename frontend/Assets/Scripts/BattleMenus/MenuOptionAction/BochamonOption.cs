using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class BochamonOption : MenuOption
{
    public Bochamon _bochamon;
    // [SerializeField] BochamonInfo _bochamonInfo;
    public override void Select(bool s)
    {
        base.Select(s);
        // _attackInfo.Setup(_move);
    }

    public void Setup(Bochamon bochamon){
        _bochamon = bochamon;
        base.Setup(bochamon.bochaName);
    }
}
