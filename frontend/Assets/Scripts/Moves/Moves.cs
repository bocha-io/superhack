using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public enum BochaType {
    Normal,
    Water,
    Fire,
    Grass
}

public class Moves : MonoBehaviour
{
    public string moveName;
    public BochaType moveType;
    public int power;
    public int speed;

    
}
