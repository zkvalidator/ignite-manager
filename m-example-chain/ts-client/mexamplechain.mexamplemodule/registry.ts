import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateEntityName } from "./types/mexamplechain/mexamplemodule/tx";
import { MsgDeleteEntityName } from "./types/mexamplechain/mexamplemodule/tx";
import { MsgUpdateEntityName } from "./types/mexamplechain/mexamplemodule/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/mexamplechain.mexamplemodule.MsgCreateEntityName", MsgCreateEntityName],
    ["/mexamplechain.mexamplemodule.MsgDeleteEntityName", MsgDeleteEntityName],
    ["/mexamplechain.mexamplemodule.MsgUpdateEntityName", MsgUpdateEntityName],
    
];

export { msgTypes }