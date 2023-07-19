            {{if eq .name "status"}}<el-form-item :label="$t('common.status')" prop="status">
                <common-status-selector class="!w-full" v-model="formData.status" @change="onChangeStatus" />
            </el-form-item>
            {{else}}<el-form-item :label="$t('{{.lowerStartCamelObject}}.itemFields.{{.name}}')" prop="{{.name}}">
                <el-input v-model="formData.{{.name}}" />
            </el-form-item>{{end}}
