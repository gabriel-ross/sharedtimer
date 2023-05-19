import React from "react";
import QuickCancelIcon from "./quickCancelIcon";
import QuickPauseIcon from "./quickPauseIcon";
import QuickStartIcon from "./quickStartIcon";
import {useNavigate} from "react-router-dom";

function TimerQuickview(props) {
    const navigate = useNavigate();
    function handleClick(timerId) {
        navigate(`/timers/${timerId}`, {state: {
            timerId: props.id
        }})
    }

    return (
        <tr onClick={() => handleClick}>
            {props.dataFields.map((key, idx) => (
                <td key={idx}>{props.rowData[key]}</td>
            ))}
            <td>
                < QuickCancelIcon timerId={props.id}/>
            </td>
            <td>
                < QuickPauseIcon timerId={props.id}/>
            </td>
            <td>
                < QuickStartIcon timerId={props.id}/>
            </td>
        </tr>
    )
}

export default TimerQuickview