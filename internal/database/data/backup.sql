PGDMP  #    
        	        |            postgres    16.3    16.3 b    ,           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            -           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            .           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            /           1262    5    postgres    DATABASE     |   CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';
    DROP DATABASE postgres;
                postgres    false            0           0    0    DATABASE postgres    COMMENT     N   COMMENT ON DATABASE postgres IS 'default administrative connection database';
                   postgres    false    4911                        2615    16397    test    SCHEMA        CREATE SCHEMA test;
    DROP SCHEMA test;
                postgres    false                        3079    16384 	   adminpack 	   EXTENSION     A   CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;
    DROP EXTENSION adminpack;
                   false            1           0    0    EXTENSION adminpack    COMMENT     M   COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';
                        false    2            �            1255    17273    log_user_registration()    FUNCTION     �   CREATE FUNCTION public.log_user_registration() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
	BEGIN
		PERFORM pg_notify('user_registration', 'User registered with ID: ' || NEW.user_id || ', Name: ' || NEW.name);
		RETURN NEW;
	END;
	$$;
 .   DROP FUNCTION public.log_user_registration();
       public          postgres    false            �            1259    17446    admins    TABLE     �   CREATE TABLE public.admins (
    admin_id integer NOT NULL,
    login character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    name text NOT NULL,
    access_level bigint,
    phone text,
    gym_id bigint
);
    DROP TABLE public.admins;
       public         heap    postgres    false            �            1259    17445    admins_admin_id_seq    SEQUENCE     �   CREATE SEQUENCE public.admins_admin_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.admins_admin_id_seq;
       public          postgres    false    225            2           0    0    admins_admin_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.admins_admin_id_seq OWNED BY public.admins.admin_id;
          public          postgres    false    224            �            1259    17500    availability    TABLE     �   CREATE TABLE public.availability (
    id integer NOT NULL,
    gym_id bigint NOT NULL,
    equipment_id bigint NOT NULL,
    amount bigint NOT NULL
);
     DROP TABLE public.availability;
       public         heap    postgres    false            �            1259    17499    availability_id_seq    SEQUENCE     �   CREATE SEQUENCE public.availability_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.availability_id_seq;
       public          postgres    false    237            3           0    0    availability_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.availability_id_seq OWNED BY public.availability.id;
          public          postgres    false    236            �            1259    17455    coaches    TABLE     �   CREATE TABLE public.coaches (
    coach_id integer NOT NULL,
    login character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    name text NOT NULL,
    age bigint,
    sex text,
    description text,
    gym_id bigint
);
    DROP TABLE public.coaches;
       public         heap    postgres    false            �            1259    17454    coaches_coach_id_seq    SEQUENCE     �   CREATE SEQUENCE public.coaches_coach_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.coaches_coach_id_seq;
       public          postgres    false    227            4           0    0    coaches_coach_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.coaches_coach_id_seq OWNED BY public.coaches.coach_id;
          public          postgres    false    226            �            1259    17507    current_events    TABLE     �   CREATE TABLE public.current_events (
    id integer NOT NULL,
    gym_id bigint NOT NULL,
    event_id bigint NOT NULL,
    "time" time(0) without time zone NOT NULL,
    data date NOT NULL
);
 "   DROP TABLE public.current_events;
       public         heap    postgres    false            �            1259    17506    current_events_id_seq    SEQUENCE     �   CREATE SEQUENCE public.current_events_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE public.current_events_id_seq;
       public          postgres    false    239            5           0    0    current_events_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE public.current_events_id_seq OWNED BY public.current_events.id;
          public          postgres    false    238            �            1259    17491    events    TABLE     �   CREATE TABLE public.events (
    event_id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    coach_id bigint NOT NULL
);
    DROP TABLE public.events;
       public         heap    postgres    false            �            1259    17490    events_event_id_seq    SEQUENCE     �   CREATE SEQUENCE public.events_event_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.events_event_id_seq;
       public          postgres    false    235            6           0    0    events_event_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.events_event_id_seq OWNED BY public.events.event_id;
          public          postgres    false    234            �            1259    17482    fitness_equipment    TABLE     �   CREATE TABLE public.fitness_equipment (
    equipment_id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL
);
 %   DROP TABLE public.fitness_equipment;
       public         heap    postgres    false            �            1259    17481 "   fitness_equipment_equipment_id_seq    SEQUENCE     �   CREATE SEQUENCE public.fitness_equipment_equipment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 9   DROP SEQUENCE public.fitness_equipment_equipment_id_seq;
       public          postgres    false    233            7           0    0 "   fitness_equipment_equipment_id_seq    SEQUENCE OWNED BY     i   ALTER SEQUENCE public.fitness_equipment_equipment_id_seq OWNED BY public.fitness_equipment.equipment_id;
          public          postgres    false    232            �            1259    17464 
   gym_passes    TABLE     �   CREATE TABLE public.gym_passes (
    gp_id integer NOT NULL,
    name text NOT NULL,
    price bigint NOT NULL,
    description text NOT NULL
);
    DROP TABLE public.gym_passes;
       public         heap    postgres    false            �            1259    17463    gym_passes_gp_id_seq    SEQUENCE     �   CREATE SEQUENCE public.gym_passes_gp_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.gym_passes_gp_id_seq;
       public          postgres    false    229            8           0    0    gym_passes_gp_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.gym_passes_gp_id_seq OWNED BY public.gym_passes.gp_id;
          public          postgres    false    228            �            1259    17473    gyms    TABLE     �   CREATE TABLE public.gyms (
    gym_id integer NOT NULL,
    name text NOT NULL,
    address text NOT NULL,
    description text NOT NULL
);
    DROP TABLE public.gyms;
       public         heap    postgres    false            �            1259    17472    gyms_gym_id_seq    SEQUENCE     �   CREATE SEQUENCE public.gyms_gym_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.gyms_gym_id_seq;
       public          postgres    false    231            9           0    0    gyms_gym_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.gyms_gym_id_seq OWNED BY public.gyms.gym_id;
          public          postgres    false    230            �            1259    17437    users    TABLE       CREATE TABLE public.users (
    user_id integer NOT NULL,
    login character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    active boolean,
    name text NOT NULL,
    age bigint,
    sex text,
    phone text,
    pass_id bigint,
    gym_id bigint
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    17436    users_user_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.users_user_id_seq;
       public          postgres    false    223            :           0    0    users_user_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;
          public          postgres    false    222            �            1259    16407    category    TABLE     P   CREATE TABLE test.category (
    title text NOT NULL,
    id bigint NOT NULL
);
    DROP TABLE test.category;
       test         heap    postgres    false    7            �            1259    16432    category_id_seq    SEQUENCE     �   ALTER TABLE test.category ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME test.category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            test          postgres    false    7    220            �            1259    16404    priority    TABLE     i   CREATE TABLE test.priority (
    title text NOT NULL,
    color text NOT NULL,
    id bigint NOT NULL
);
    DROP TABLE test.priority;
       test         heap    postgres    false    7            �            1259    16398    task    TABLE     �   CREATE TABLE test.task (
    title text NOT NULL,
    completed numeric NOT NULL,
    task_date time without time zone,
    id bigint NOT NULL,
    category_id bigint
);
    DROP TABLE test.task;
       test         heap    postgres    false    7            ;           0    0 
   TABLE task    COMMENT     A   COMMENT ON TABLE test.task IS 'тестовая таблица';
          test          postgres    false    217            �            1259    16401 	   user_data    TABLE     �   CREATE TABLE test.user_data (
    email text NOT NULL,
    password text NOT NULL,
    username text NOT NULL,
    id bigint NOT NULL
);
    DROP TABLE test.user_data;
       test         heap    postgres    false    7            W           2604    17559    admins admin_id    DEFAULT     r   ALTER TABLE ONLY public.admins ALTER COLUMN admin_id SET DEFAULT nextval('public.admins_admin_id_seq'::regclass);
 >   ALTER TABLE public.admins ALTER COLUMN admin_id DROP DEFAULT;
       public          postgres    false    225    224    225            ]           2604    17560    availability id    DEFAULT     r   ALTER TABLE ONLY public.availability ALTER COLUMN id SET DEFAULT nextval('public.availability_id_seq'::regclass);
 >   ALTER TABLE public.availability ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    236    237    237            X           2604    17561    coaches coach_id    DEFAULT     t   ALTER TABLE ONLY public.coaches ALTER COLUMN coach_id SET DEFAULT nextval('public.coaches_coach_id_seq'::regclass);
 ?   ALTER TABLE public.coaches ALTER COLUMN coach_id DROP DEFAULT;
       public          postgres    false    226    227    227            ^           2604    17562    current_events id    DEFAULT     v   ALTER TABLE ONLY public.current_events ALTER COLUMN id SET DEFAULT nextval('public.current_events_id_seq'::regclass);
 @   ALTER TABLE public.current_events ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    238    239    239            \           2604    17563    events event_id    DEFAULT     r   ALTER TABLE ONLY public.events ALTER COLUMN event_id SET DEFAULT nextval('public.events_event_id_seq'::regclass);
 >   ALTER TABLE public.events ALTER COLUMN event_id DROP DEFAULT;
       public          postgres    false    234    235    235            [           2604    17564    fitness_equipment equipment_id    DEFAULT     �   ALTER TABLE ONLY public.fitness_equipment ALTER COLUMN equipment_id SET DEFAULT nextval('public.fitness_equipment_equipment_id_seq'::regclass);
 M   ALTER TABLE public.fitness_equipment ALTER COLUMN equipment_id DROP DEFAULT;
       public          postgres    false    232    233    233            Y           2604    17565    gym_passes gp_id    DEFAULT     t   ALTER TABLE ONLY public.gym_passes ALTER COLUMN gp_id SET DEFAULT nextval('public.gym_passes_gp_id_seq'::regclass);
 ?   ALTER TABLE public.gym_passes ALTER COLUMN gp_id DROP DEFAULT;
       public          postgres    false    228    229    229            Z           2604    17566    gyms gym_id    DEFAULT     j   ALTER TABLE ONLY public.gyms ALTER COLUMN gym_id SET DEFAULT nextval('public.gyms_gym_id_seq'::regclass);
 :   ALTER TABLE public.gyms ALTER COLUMN gym_id DROP DEFAULT;
       public          postgres    false    230    231    231            V           2604    17567    users user_id    DEFAULT     n   ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);
 <   ALTER TABLE public.users ALTER COLUMN user_id DROP DEFAULT;
       public          postgres    false    223    222    223                      0    17446    admins 
   TABLE DATA           ^   COPY public.admins (admin_id, login, password, name, access_level, phone, gym_id) FROM stdin;
    public          postgres    false    225   o       '          0    17500    availability 
   TABLE DATA           H   COPY public.availability (id, gym_id, equipment_id, amount) FROM stdin;
    public          postgres    false    237   �o                 0    17455    coaches 
   TABLE DATA           a   COPY public.coaches (coach_id, login, password, name, age, sex, description, gym_id) FROM stdin;
    public          postgres    false    227   �o       )          0    17507    current_events 
   TABLE DATA           L   COPY public.current_events (id, gym_id, event_id, "time", data) FROM stdin;
    public          postgres    false    239   �o       %          0    17491    events 
   TABLE DATA           G   COPY public.events (event_id, name, description, coach_id) FROM stdin;
    public          postgres    false    235   �o       #          0    17482    fitness_equipment 
   TABLE DATA           L   COPY public.fitness_equipment (equipment_id, name, description) FROM stdin;
    public          postgres    false    233   �o                 0    17464 
   gym_passes 
   TABLE DATA           E   COPY public.gym_passes (gp_id, name, price, description) FROM stdin;
    public          postgres    false    229   �r       !          0    17473    gyms 
   TABLE DATA           B   COPY public.gyms (gym_id, name, address, description) FROM stdin;
    public          postgres    false    231   �s                 0    17437    users 
   TABLE DATA           i   COPY public.users (user_id, login, password, active, name, age, sex, phone, pass_id, gym_id) FROM stdin;
    public          postgres    false    223   cu                 0    16407    category 
   TABLE DATA           +   COPY test.category (title, id) FROM stdin;
    test          postgres    false    220   �v                 0    16404    priority 
   TABLE DATA           2   COPY test.priority (title, color, id) FROM stdin;
    test          postgres    false    219   �v                 0    16398    task 
   TABLE DATA           J   COPY test.task (title, completed, task_date, id, category_id) FROM stdin;
    test          postgres    false    217   �v                 0    16401 	   user_data 
   TABLE DATA           @   COPY test.user_data (email, password, username, id) FROM stdin;
    test          postgres    false    218   ,w       <           0    0    admins_admin_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.admins_admin_id_seq', 2, true);
          public          postgres    false    224            =           0    0    availability_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.availability_id_seq', 1, false);
          public          postgres    false    236            >           0    0    coaches_coach_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.coaches_coach_id_seq', 1, false);
          public          postgres    false    226            ?           0    0    current_events_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.current_events_id_seq', 1, false);
          public          postgres    false    238            @           0    0    events_event_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.events_event_id_seq', 1, false);
          public          postgres    false    234            A           0    0 "   fitness_equipment_equipment_id_seq    SEQUENCE SET     Q   SELECT pg_catalog.setval('public.fitness_equipment_equipment_id_seq', 1, false);
          public          postgres    false    232            B           0    0    gym_passes_gp_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.gym_passes_gp_id_seq', 1, false);
          public          postgres    false    228            C           0    0    gyms_gym_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.gyms_gym_id_seq', 1, false);
          public          postgres    false    230            D           0    0    users_user_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.users_user_id_seq', 6, true);
          public          postgres    false    222            E           0    0    category_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('test.category_id_seq', 12, true);
          test          postgres    false    221            j           2606    17453    admins admins_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_pkey PRIMARY KEY (admin_id);
 <   ALTER TABLE ONLY public.admins DROP CONSTRAINT admins_pkey;
       public            postgres    false    225            v           2606    17505    availability availability_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.availability
    ADD CONSTRAINT availability_pkey PRIMARY KEY (id);
 H   ALTER TABLE ONLY public.availability DROP CONSTRAINT availability_pkey;
       public            postgres    false    237            l           2606    17462    coaches coaches_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.coaches
    ADD CONSTRAINT coaches_pkey PRIMARY KEY (coach_id);
 >   ALTER TABLE ONLY public.coaches DROP CONSTRAINT coaches_pkey;
       public            postgres    false    227            x           2606    17512 "   current_events current_events_pkey 
   CONSTRAINT     `   ALTER TABLE ONLY public.current_events
    ADD CONSTRAINT current_events_pkey PRIMARY KEY (id);
 L   ALTER TABLE ONLY public.current_events DROP CONSTRAINT current_events_pkey;
       public            postgres    false    239            t           2606    17498    events events_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (event_id);
 <   ALTER TABLE ONLY public.events DROP CONSTRAINT events_pkey;
       public            postgres    false    235            r           2606    17489 (   fitness_equipment fitness_equipment_pkey 
   CONSTRAINT     p   ALTER TABLE ONLY public.fitness_equipment
    ADD CONSTRAINT fitness_equipment_pkey PRIMARY KEY (equipment_id);
 R   ALTER TABLE ONLY public.fitness_equipment DROP CONSTRAINT fitness_equipment_pkey;
       public            postgres    false    233            n           2606    17471    gym_passes gym_passes_pkey 
   CONSTRAINT     [   ALTER TABLE ONLY public.gym_passes
    ADD CONSTRAINT gym_passes_pkey PRIMARY KEY (gp_id);
 D   ALTER TABLE ONLY public.gym_passes DROP CONSTRAINT gym_passes_pkey;
       public            postgres    false    229            p           2606    17480    gyms gyms_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.gyms
    ADD CONSTRAINT gyms_pkey PRIMARY KEY (gym_id);
 8   ALTER TABLE ONLY public.gyms DROP CONSTRAINT gyms_pkey;
       public            postgres    false    231            h           2606    17444    users users_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    223            f           2606    16425    category category_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY test.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY test.category DROP CONSTRAINT category_pkey;
       test            postgres    false    220            d           2606    16419    priority priority_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY test.priority
    ADD CONSTRAINT priority_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY test.priority DROP CONSTRAINT priority_pkey;
       test            postgres    false    219            `           2606    16421    task task_pkey 
   CONSTRAINT     J   ALTER TABLE ONLY test.task
    ADD CONSTRAINT task_pkey PRIMARY KEY (id);
 6   ALTER TABLE ONLY test.task DROP CONSTRAINT task_pkey;
       test            postgres    false    217            b           2606    16413    user_data user_data_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY test.user_data
    ADD CONSTRAINT user_data_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY test.user_data DROP CONSTRAINT user_data_pkey;
       test            postgres    false    218            �           2620    17568    users user_registration_trigger    TRIGGER     �   CREATE TRIGGER user_registration_trigger AFTER INSERT ON public.users FOR EACH ROW EXECUTE FUNCTION public.log_user_registration();
 8   DROP TRIGGER user_registration_trigger ON public.users;
       public          postgres    false    240    223            |           2606    17528    admins admins_gym_id_foreign    FK CONSTRAINT     }   ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_gym_id_foreign FOREIGN KEY (gym_id) REFERENCES public.gyms(gym_id);
 F   ALTER TABLE ONLY public.admins DROP CONSTRAINT admins_gym_id_foreign;
       public          postgres    false    4720    231    225                       2606    17553 .   availability availability_equipment_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY public.availability
    ADD CONSTRAINT availability_equipment_id_foreign FOREIGN KEY (equipment_id) REFERENCES public.fitness_equipment(equipment_id);
 X   ALTER TABLE ONLY public.availability DROP CONSTRAINT availability_equipment_id_foreign;
       public          postgres    false    233    237    4722            �           2606    17523 (   availability availability_gym_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY public.availability
    ADD CONSTRAINT availability_gym_id_foreign FOREIGN KEY (gym_id) REFERENCES public.gyms(gym_id);
 R   ALTER TABLE ONLY public.availability DROP CONSTRAINT availability_gym_id_foreign;
       public          postgres    false    237    231    4720            }           2606    17518    coaches coaches_gym_id_foreign    FK CONSTRAINT        ALTER TABLE ONLY public.coaches
    ADD CONSTRAINT coaches_gym_id_foreign FOREIGN KEY (gym_id) REFERENCES public.gyms(gym_id);
 H   ALTER TABLE ONLY public.coaches DROP CONSTRAINT coaches_gym_id_foreign;
       public          postgres    false    231    4720    227            �           2606    17548 .   current_events current_events_event_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY public.current_events
    ADD CONSTRAINT current_events_event_id_foreign FOREIGN KEY (event_id) REFERENCES public.events(event_id);
 X   ALTER TABLE ONLY public.current_events DROP CONSTRAINT current_events_event_id_foreign;
       public          postgres    false    4724    235    239            �           2606    17543 ,   current_events current_events_gym_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY public.current_events
    ADD CONSTRAINT current_events_gym_id_foreign FOREIGN KEY (gym_id) REFERENCES public.gyms(gym_id);
 V   ALTER TABLE ONLY public.current_events DROP CONSTRAINT current_events_gym_id_foreign;
       public          postgres    false    231    4720    239            ~           2606    17513    events events_coach_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_coach_id_foreign FOREIGN KEY (coach_id) REFERENCES public.coaches(coach_id);
 H   ALTER TABLE ONLY public.events DROP CONSTRAINT events_coach_id_foreign;
       public          postgres    false    235    227    4716            z           2606    17538    users users_gym_id_foreign    FK CONSTRAINT     {   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_gym_id_foreign FOREIGN KEY (gym_id) REFERENCES public.gyms(gym_id);
 D   ALTER TABLE ONLY public.users DROP CONSTRAINT users_gym_id_foreign;
       public          postgres    false    231    223    4720            {           2606    17533    users users_pass_id_foreign    FK CONSTRAINT     �   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pass_id_foreign FOREIGN KEY (pass_id) REFERENCES public.gym_passes(gp_id);
 E   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pass_id_foreign;
       public          postgres    false    4718    229    223            y           2606    16433    task category_fkey    FK CONSTRAINT     ~   ALTER TABLE ONLY test.task
    ADD CONSTRAINT category_fkey FOREIGN KEY (category_id) REFERENCES test.category(id) NOT VALID;
 :   ALTER TABLE ONLY test.task DROP CONSTRAINT category_fkey;
       test          postgres    false    220    4710    217               j   x�3�LL���342�T1JT14P	*Mv312�K�4�(I��(2���M��56��4Ks2�qw���.r��L��.�4�0�¦�/l���;9c� �+F��� }�#T      '      x������ � �            x������ � �      )      x������ � �      %      x������ � �      #   �  x��UKnA]w�b.�ɏ���O�#9QH,P 6l�Ǝ'���
�7���������Q<��{�^��F>J*R�Tw�Lj�����#�s7��
j5���~����VK��?�R�F)���������࢐��7�N.B��M*��N�	�k掐�9{h��Zf�"�%u=#��O�R�Z{ǈE�.b))W��Z�LQN����i%����_@g�YϘ��n;E))"]��,���Z��̍��x��n$�5Xdk�#5p=7\9���w,Ҟ��/�b�U�x�����&d
�D����W��^��Z�D���3
���.��Ϟz��U�JF�q;׍;G�
�nE�m���nό|%p��[��F�
)��S$l4������sM�0��F�tj��͍R[��qL�!���\J;�^�H�޷�X�77���/���;�(�o߼c���@5�L�i:��`7v��Q��e�&�ST��%�?Z��l�F��J���H��%��s�p��db���/q�g�5�1�-~�Jt��?��R�hD@�G}MqW/uk�$6>�,�.U�.:}@X�7�.ݰ� �}��Ƙ���-�33���9WR���X|r����ÄL�dςr�(�(?Y~ 19g�Χ�wj}�g�`��8lpm�5���87��"XΝ��~���Q������}��Z���M�         �   x���M�0���)zÏ��0Qb0�D�.<BA�H�¼�^�;m�7_g���p�QM%��2�h�1���k�\%Q)\�̩�hX�-ep�G[���ѢG��ٹ���-�au�p�gཨ��3q�^P�|�ǃK3�T�T���w�5x�=�X�V8�f�WE��+�̚P�V�|�bj�L#�B�0?�/����7+���]+x      !   �  x��S�J�@��{�{�Z��>����BĨ���O����M����9{i�X)�۽��ٙ�0�;r���R�#|`�R&Z<�@���LQ�-^�E%!�HY�$��a�|�TBVn���1VV�;�r�Ǖ���7�*��MĪr+S�P��i&�(��GgG|.�u`�Caq�6��o�暒K��Hr����"ߐ��:��D1�:/S�: ~��j;�O̘�s.M3�x��c���'����L+=\*�*�I�7��QJr��d� �镒x��v��+=��ʝy��uVs�~h��g�>�"�~���	�m#,����38C�Z�{�T�~[R=i���8�3�6t���NU���[�ML'�I���:>s4Ybs�y�~9�ҬC��p� �n^��c��OnĤ��;�����_�|T|ʦ��������9�c~ ��8          $  x�3�,-N-�,H,.���0�¦�/l���;A"H�ˈ31%73/>?�2����Ș�cΜ̪�D�V=�j����L�� ��b��bh�����Wd^ᝫ�����X��T�aR�V�g��W��Unf�f蛜k�J�9�`kLL`֤�闧��dZ�g8�9������&e���D���8�f䕹[&:��'�Xdx�ƌ�F�0[�J��L�L�3M=J+�+J��#�s�2}�Mr3�ҜLs���<��\+2�"̓��,����������qqq �ƞ         *   x����bÅ}�/콰��А���[.6"q��qqq "�            x������ � �         .   x��0�¾�.l�دpa��� ��6pr��	CC�=... �M�             x������p�yaÅM0��14����� _�      