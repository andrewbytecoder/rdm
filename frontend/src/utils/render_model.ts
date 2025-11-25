import {h} from "vue";
import {NIcon} from "naive-ui";


// 图标渲染函数，对图标进行渲染
export const renderIcon = (icon: any) => {
    return () => {
        return h(NIcon, null, {
            default: () => h(icon),
        })
    }
}
