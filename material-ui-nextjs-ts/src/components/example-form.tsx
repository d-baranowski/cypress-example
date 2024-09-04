import React, {useState} from "react";
import {Controller, useFieldArray, useForm} from "react-hook-form";
import {z} from "zod";
import {zodResolver} from "@hookform/resolvers/zod";
import {
  AutocompleteElement,
  CheckboxElement,
  DateTimePickerElement,
  RadioButtonGroup,
  SliderElement,
  TextFieldElement,
} from "react-hook-form-mui";
import {Box, Button, IconButton, Tab} from "@mui/material";
import DeleteIcon from '@mui/icons-material/Delete';
import {TabContext, TabList, TabPanel} from '@mui/lab';
import {saveFormModel} from "../../gen/example/v1/form-FormModelService_connectquery";
import {useMutation, useTransport} from "@connectrpc/connect-query";
import toast from "react-hot-toast";
import {createPromiseClient} from "@connectrpc/connect";
import {FormModelService} from "../../gen/example/v1/form_connect";

interface Props {
  id?: bigint;
  onCancel: () => void;
  afterSave: () => void;
}

const ExampleForm: React.FC<Props> = function ExampleForm(props) {
  const [tabValue, setTabValue] = useState('1');
  const handleTabChange = (event: React.SyntheticEvent, newValue: string) => {
    setTabValue(newValue);
  };

  const {afterSave} = props;
  const {mutateAsync} = useMutation(saveFormModel);

  const {
    control,
    handleSubmit,
    formState: {errors},
    reset,
  } = useForm({
    resolver: zodResolver(schema),
    defaultValues: {
      name: "",
      age: null,
      dateTime: null,
      skills: [],
      city: null,
      range: [20, 50],
      termsAccepted: false,
      satisfactionLevel: 5,
      gender: 'male',
      todos: [{task: ""}],
    },
  });

  const t = useTransport();
  const client = createPromiseClient(FormModelService, t)
  React.useEffect(() => {
    if (props.id) {
      toast.promise(client.getFormModel({id: String(props.id)}).then(res => {
        return reset({...res.formModel as any});
      }), {
        loading: 'Loading',
        success: 'Loaded the form',
        error: 'Error when loading',
      })
    }
  }, [props.id])

  const {fields, append, remove} = useFieldArray({
    control,
    name: "todos",
  });

  const onSubmit = async (data: any) => {
    toast.promise(mutateAsync({
      formModel: {
        id: props.id,
        ...data,
        dateTime: data.dateTime.toISOString(),
      },
    }), {
      loading: 'Loading',
      success: 'Saved the form',
      error: 'Error when saving',
    });
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <TabContext value={tabValue}>
        <Box sx={{borderBottom: 1, borderColor: 'divider'}}>
          <TabList onChange={handleTabChange} aria-label="Form tabs">
            <Tab label="Personal Info" value="1"/>
            <Tab label="Preferences" value="2"/>
            <Tab label="Tasks" value="3"/>
          </TabList>
        </Box>

        {/* First Tab: Personal Info */}
        <TabPanel value="1">
          <TextFieldElement
            name="name"
            label="Name"
            required
            control={control}
            error={!!errors.name}
            helperText={errors.name?.message}
          />

          <TextFieldElement
            name="age"
            label="Age"
            type="number"
            required
            control={control}
            error={!!errors.age}
            helperText={errors.age?.message}
          />

          <DateTimePickerElement
            name="dateTime"
            label="Select Date & Time"
            required
            control={control}
          />

          <RadioButtonGroup
            name="gender"
            label="Select Gender"
            options={genderOptions}
            control={control}
            required
          />
        </TabPanel>

        {/* Second Tab: Preferences */}
        <TabPanel value="2">
          <AutocompleteElement
            name="skills"
            label="Select Skills"
            multiple
            options={skillsOptions}
            required
            control={control}
          />

          <AutocompleteElement
            name="city"
            label="Select City"
            options={cityOptions}
            required
            control={control}
          />

          <Controller
            name="range"
            control={control}
            render={({field}) => (
              <Box sx={{mt: 2}}>
                <SliderElement
                  {...field}
                  label="Select Range"
                  min={0}
                  max={100}
                  step={1}
                  valueLabelDisplay="auto"
                  control={control}
                />
                {errors.range && (
                  <p style={{color: "red"}}>{errors.range?.message}</p>
                )}
              </Box>
            )}
          />

          <SliderElement
            name="satisfactionLevel"
            label="Satisfaction Level"
            min={0}
            max={10}
            step={1}
            control={control}
          />

          <CheckboxElement
            name="termsAccepted"
            label="I accept the terms and conditions"
            required
            control={control}
          />
        </TabPanel>

        {/* Third Tab: Tasks */}
        <TabPanel value="3">
          <Box sx={{mt: 2}}>
            <h4>Todo List</h4>
            {fields.map((item, index) => (
              <Box key={item.id} sx={{display: "flex", alignItems: "center"}}>
                <TextFieldElement
                  name={`todos.${index}.task`}
                  label={`Task ${index + 1}`}
                  required
                  control={control}
                  sx={{flexGrow: 1}}
                  error={!!errors.todos?.[index]?.task}
                  helperText={errors.todos?.[index]?.task?.message}
                />
                <IconButton
                  aria-label="delete"
                  onClick={() => remove(index)}
                  color="error"
                >
                  <DeleteIcon/>
                </IconButton>
              </Box>
            ))}
            <Button onClick={() => append({task: ""})} variant="contained" sx={{mt: 2}}>
              Add Task
            </Button>
          </Box>
        </TabPanel>
      </TabContext>

      {/* Submit button */}
      <Button type="submit" variant="contained" sx={{mt: 4}}>
        Submit
      </Button>
    </form>
  );
}

// Define the Zod schema
const schema = z.object({
  name: z.string().nonempty("Name is required"),
  age: z.number().min(1, "Age must be a positive number"),
  dateTime: z.date(),
  skills: z.array(z.string()).nonempty("Please select at least one skill"),
  city: z.string().nonempty("Please select a city"),
  range: z.array(z.number()).length(2, "Range should have two values"),
  termsAccepted: z.boolean().refine(val => val === true, "You must accept the terms"),
  satisfactionLevel: z.number().min(0).max(10),
  gender: z.enum(["male", "female", "other"]),
  todos: z.array(
    z.object({
      task: z.string().nonempty("Task is required"),
    })
  ),
});

const skillsOptions = ["JavaScript", "TypeScript", "React", "Node.js"];
const cityOptions = ["New York", "San Francisco", "Chicago", "Los Angeles"];
const genderOptions = [
  {id: 'male', label: 'Male'},
  {id: 'female', label: 'Female'},
  {id: 'other', label: 'Other'},
];

export default ExampleForm;
