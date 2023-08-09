// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {MonType, ElementType, AttackType} from "../codegen/Types.sol";

library LibDefaults {
    function getHp(MonType mon) internal pure returns (int32) {
        if (mon == MonType.Cobarett) {
            return 350;
        }
        if (mon == MonType.Flarezael) {
            return 450;
        }
        if (mon == MonType.Firomenis) {
            return 300;
        }
        if (mon == MonType.Baobaffe) {
            return 500;
        }
        if (mon == MonType.Howliage) {
            return 350;
        }
        if (mon == MonType.Sunnydra) {
            return 420;
        }
        if (mon == MonType.Tobishimi) {
            return 400;
        }
        if (mon == MonType.Mobiusk) {
            return 300;
        }
        if (mon == MonType.Ramon) {
            return 450;
        }
        return 0;
    }

    function getMonSpeedAndType(MonType mon) internal pure returns (int32, ElementType) {
        if (mon == MonType.Cobarett) {
            return (100, ElementType.Fire);
        }
        if (mon == MonType.Flarezael) {
            return (60, ElementType.Fire);
        }
        if (mon == MonType.Firomenis) {
            return (120, ElementType.Fire);
        }
        if (mon == MonType.Baobaffe) {
            return (50, ElementType.Grass);
        }
        if (mon == MonType.Howliage) {
            return (100, ElementType.Grass);
        }
        if (mon == MonType.Sunnydra) {
            return (60, ElementType.Grass);
        }
        if (mon == MonType.Tobishimi) {
            return (80, ElementType.Water);
        }
        if (mon == MonType.Mobiusk) {
            return (120, ElementType.Water);
        }
        if (mon == MonType.Ramon) {
            return (70, ElementType.Water);
        }
        return (0, ElementType.Normal);
    }

    function getAttack(MonType mon, uint8 pos) internal pure returns (int32, int32, ElementType) {
        if (mon == MonType.Cobarett) {
            if (pos == 0) {
                return getAttackValues(AttackType.LavaPlume);
            }
            if (pos == 1) {
                return getAttackValues(AttackType.FireLash);
            }
            if (pos == 2) {
                return getAttackValues(AttackType.Bite);
            }
            if (pos == 3) {
                return getAttackValues(AttackType.Tackle);
            }
        }

        if (mon == MonType.Flarezael) {
            if (pos == 0) {
                return getAttackValues(AttackType.FieryDance);
            }
            if (pos == 1) {
                return getAttackValues(AttackType.FireLash);
            }
            if (pos == 2) {
                return getAttackValues(AttackType.CrushGrip);
            }
            if (pos == 3) {
                return getAttackValues(AttackType.Flail);
            }
        }

        if (mon == MonType.Firomenis) {
            if (pos == 0) {
                return getAttackValues(AttackType.FieryDance);
            }
            if (pos == 1) {
                return getAttackValues(AttackType.LavaPlume);
            }
            if (pos == 2) {
                return getAttackValues(AttackType.Tackle);
            }
            if (pos == 3) {
                return getAttackValues(AttackType.Flail);
            }
        }

        if (mon == MonType.Baobaffe) {
            if (pos == 0) {
                return getAttackValues(AttackType.LeafTornado);
            }
            if (pos == 1) {
                return getAttackValues(AttackType.RazorLeaf);
            }
            if (pos == 2) {
                return getAttackValues(AttackType.Tackle);
            }
            if (pos == 3) {
                return getAttackValues(AttackType.Flail);
            }
        }

        if (mon == MonType.Howliage) {
            if (pos == 0) {
                return getAttackValues(AttackType.SolarBlade);
            }
            if (pos == 1) {
                return getAttackValues(AttackType.RazorLeaf);
            }
            if (pos == 2) {
                return getAttackValues(AttackType.Tackle);
            }
            if (pos == 3) {
                return getAttackValues(AttackType.Bite);
            }
        }

        if (mon == MonType.Sunnydra) {
            if (pos == 0) {
                return getAttackValues(AttackType.SolarBlade);
            }
            if (pos == 1) {
                return getAttackValues(AttackType.LeafTornado);
            }
            if (pos == 2) {
                return getAttackValues(AttackType.Flail);
            }
            if (pos == 3) {
                return getAttackValues(AttackType.Bite);
            }
        }

        if (mon == MonType.Tobishimi) {
            if (pos == 0) {
                return getAttackValues(AttackType.AquaTail);
            }
            if (pos == 1) {
                return getAttackValues(AttackType.BubbleBeam);
            }
            if (pos == 2) {
                return getAttackValues(AttackType.Flail);
            }
            if (pos == 3) {
                return getAttackValues(AttackType.Bite);
            }
        }

        if (mon == MonType.Mobiusk) {
            if (pos == 0) {
                return getAttackValues(AttackType.AquaTail);
            }
            if (pos == 1) {
                return getAttackValues(AttackType.HydroVortex);
            }
            if (pos == 2) {
                return getAttackValues(AttackType.Tackle);
            }
            if (pos == 3) {
                return getAttackValues(AttackType.Bite);
            }
        }

        if (mon == MonType.Ramon) {
            if (pos == 0) {
                return getAttackValues(AttackType.BubbleBeam);
            }
            if (pos == 1) {
                return getAttackValues(AttackType.HydroVortex);
            }
            if (pos == 2) {
                return getAttackValues(AttackType.Tackle);
            }
            if (pos == 3) {
                return getAttackValues(AttackType.Flail);
            }
        }
        return (0, 0, ElementType.Normal);
    }

    function getAttackValues(AttackType attack) internal pure returns (int32, int32, ElementType) {
        // Fire
        if (attack == AttackType.LavaPlume) {
            return (80, 30, ElementType.Fire);
        }
        if (attack == AttackType.FireLash) {
            return (50, 50, ElementType.Fire);
        }
        if (attack == AttackType.FieryDance) {
            return (110, 20, ElementType.Fire);
        }

        // Grass
        if (attack == AttackType.LeafTornado) {
            return (120, 30, ElementType.Grass);
        }
        if (attack == AttackType.RazorLeaf) {
            return (50, 50, ElementType.Grass);
        }
        if (attack == AttackType.SolarBlade) {
            return (80, 40, ElementType.Grass);
        }

        // Water
        if (attack == AttackType.AquaTail) {
            return (60, 50, ElementType.Water);
        }
        if (attack == AttackType.BubbleBeam) {
            return (120, 30, ElementType.Water);
        }
        if (attack == AttackType.HydroVortex) {
            return (90, 60, ElementType.Water);
        }

        // Normal
        if (attack == AttackType.Bite) {
            return (30, 80, ElementType.Normal);
        }
        if (attack == AttackType.Tackle) {
            return (40, 60, ElementType.Normal);
        }
        if (attack == AttackType.CrushGrip) {
            return (80, 30, ElementType.Normal);
        }
        if (attack == AttackType.Flail) {
            return (20, 100, ElementType.Normal);
        }
        return (0, 0, ElementType.Normal);
    }

    function getAttackDamage(int32 dmg, ElementType attack, ElementType mon) internal pure returns (int32) {
        if (attack == ElementType.Fire) {
            if (mon == ElementType.Grass) {
                return dmg * 2;
            }
            if (mon == ElementType.Water) {
                return dmg / 2;
            }
        } else if (attack == ElementType.Grass) {
            if (mon == ElementType.Water) {
                return dmg * 2;
            }
            if (mon == ElementType.Fire) {
                return dmg / 2;
            }
        } else if (attack == ElementType.Water) {
            if (mon == ElementType.Fire) {
                return dmg * 2;
            }
            if (mon == ElementType.Grass) {
                return dmg / 2;
            }
        }
        return dmg;
    }
}
