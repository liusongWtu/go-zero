        {{if .hasSuffixTime}}{
            label: $t('{{.lowerStartCamelObject}}.itemFields.{{.name}}'),
            prop: '{{.name}}',
            minWidth: 180,
            renderHeader,
            cellRenderer: createTimestampRender('{{.name}}')
        }{{else if .hasSuffixStatus}}{
            label: $t('{{.lowerStartCamelObject}}.itemFields.{{.name}}'),
            prop: '{{.name}}',
            minWidth: 130,
            renderHeader,
            cellRenderer: statusRender
        }{{else}}{
            label: $t('{{.lowerStartCamelObject}}.itemFields.{{.name}}'),
            prop: '{{.name}}',
            minWidth: 100,
            renderHeader,
            formatter: createNormalFormatter('{{.name}}')
        }{{end}},